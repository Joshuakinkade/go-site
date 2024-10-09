import requests
import sys

def main():
    action = sys.argv[1]

    if action == 'list-posts':
        list_posts()
    elif action == 'create-post':
        create_post()
    else:
        print("Action not available")

def list_posts():
    r = requests.get('http://localhost:8080/api/v1/posts')
    print(r.text)

def create_post():
    # Define Post
    post = {
        'title': 'Another Post',
        'body': 'Here we go again!'
    }

    # Send Request
    r = requests.post('http://localhost:8080/api/v1/posts', json=post)

    # Check Result
    print(r.text)

if __name__ == '__main__':
    main()
