# this is a simple script to serve as an e2e test for the project
# unless you have python installed, run this inside the docker container
import os
import requests

protocol = os.getenv('PROTOCOL', 'http')
host = os.getenv('HOST', 'localhost')
port = os.getenv('PORT', '8080')
debug = os.getenv('DEBUG', True)


def heath_check():
    url = f'{protocol}://{host}:{port}/health'
    response = requests.get(url)

    if debug:
        print(f'Health check response: {response.status_code}')
        print(f'Health check response: {response.text}')

    assert response.status_code == 200


def create_user():
    url = f'{protocol}://{host}:{port}/user/create'
    response = requests.post(url, json=None)

    if debug:
        print(f'Create user response: {response.status_code}')
        print(f'Create user response: {response.text}')

    assert response.status_code == 201

    # parse the response
    response_json = response.json()
    return response_json['id'], response_json['email'], response_json['password']


def login_user(user_email, user_password):
    url = f'{protocol}://{host}:{port}/login'
    response = requests.post(url, json={'email': user_email, 'password': user_password})

    if debug:
        print(f'Login response: {response.status_code}')
        print(f'Login response: {response.text}')

    assert response.status_code == 200

    # parse the response
    response_json = response.json()
    return response_json['token']


def swipe(token, to_id, preference, assert_func=None):
    url = f'{protocol}://{host}:{port}/swipe'
    headers = {'Authorization': f'Bearer {token}'}
    response = requests.post(url, json={'userId': to_id, 'preference': preference}, headers=headers)

    if debug:
        print(f'Swipe response: {response.status_code}')
        print(f'Swipe response: {response.text}')

    assert response.status_code == 200

    if assert_func:
        assert_func(response.json())


def discover(token, query_params, assert_func=None):
    url = f'{protocol}://{host}:{port}/discover'
    headers = {'Authorization': f'Bearer {token}'}
    response = requests.get(url, params=query_params, headers=headers)

    if debug:
        print(f'Discover response: {response.status_code}')
        print(f'Discover response: {response.text}')

    assert response.status_code == 200

    if assert_func:
        assert_func(response.json())


# run the tests
if __name__ == '__main__':
    print('Running health check')
    heath_check()

    print('Creating users')
    user_one_id, user_one_email, user_one_password = create_user()
    user_two_id, user_two_email, user_two_password = create_user()

    for i in range(2, 10):
        print(f'Creating user {i}')
        create_user()

    user_one_token = login_user(user_one_email, user_one_password)
    user_two_token = login_user(user_two_email, user_two_password)

    print('Swiping')

    def assert_swipe_no_match(response):
        assert response['matched'] is False

    def assert_swipe_match(response):
        assert response['matched'] is True
        assert response['matchId'] is not None

    swipe(user_one_token, user_two_id, 'YES', assert_swipe_no_match)
    swipe(user_two_token, user_one_id, 'YES', assert_swipe_match)

    print('Discovering')

    def assert_discover_no_users_swiped(response):
        for user in response:
            assert user['id'] != user_one_id
            assert user['id'] != user_two_id

    discover(user_one_token, {}, assert_discover_no_users_swiped)

    # assert filters
    def assert_discover_filter(response):
        for user in response:
            assert user['age'] >= 20
            assert user['age'] <= 50
            assert user['gender'] == 'F'

    discover(user_one_token, {"min_age": 20, "max_age": 50, "gender": "F"}, assert_discover_filter)

    # assert sort
    def assert_discover_sort(response):
        last_distance = None
        for user in response:
            if last_distance:
                assert user['distanceFromMe'] >= last_distance
            last_distance = user['distanceFromMe']

    discover(user_one_token, {"sort_distance": "true"}, assert_discover_sort)

    print('All tests passed')
