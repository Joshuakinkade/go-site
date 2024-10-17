import requests
import sys

def main():
    action = sys.argv[1]

    if action == 'list-posts':
        list_posts()
    elif action == 'get-post':
        get_post()
    elif action == 'create-post':
        create_post()
    else:
        print("Action not available")

def list_posts():
    r = requests.get('http://localhost:8080/api/v1/posts')
    print(r.text)

def create_post():
    # Defin Post
    post = {
        'title': 'Fourth Post',
        'body': '## A Heading\nAnd some body text'
    }

    # Send Request
    r = requests.post('http://localhost:8080/api/v1/posts', json=post)

    # Check Result
    print(r.text)

def get_post():
    r = requests.get('http://localhost:8080/api/v1/posts/fourth-post')
    print(r.text)

if __name__ == '__main__':
    main()
