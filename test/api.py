import requests
import sys

def main():
    action = sys.argv[1]

    if action == 'list-posts':
        list_posts()
    elif action == 'get-post':
        if len(sys.argv) > 2:
            get_post(sys.argv[2])
        else:
            get_post()
    elif action == 'create-post':
        create_post()
    elif action == 'update-post':
        update_post()
    else:
        print("Action not available")

def list_posts():
    r = requests.get('http://localhost:8080/api/v1/posts')
    posts = r.json()
    for post in posts:
        print(f'{post['title']}: {post['slug']}')

def create_post():
    # Define Post
    post = {
        'title': 'Fourth Post',
        'body': '## A Heading\nAnd some body text'
    }

    # Send Request
    r = requests.post('http://localhost:8080/api/v1/posts', json=post)

    # Check Result
    print(r.text)
    print(r.status_code)

def update_post():
    updates = {
        'body': "## Update\n\nThis post has been modified from it's original text. A third time."
    }
    slug = "fourth-post"

    r = requests.patch(f'http://localhost:8080/api/v1/posts/{slug}', json=updates)
    print(r.text)

def get_post(title='fourth-post'):
    r = requests.get(f'http://localhost:8080/api/v1/posts/{title}')
    post = r.json()
    print(post['title'])
    print(post['body'])

if __name__ == '__main__':
    main()
