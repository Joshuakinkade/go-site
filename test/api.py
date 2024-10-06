import requests

def createPost():
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
    createPost()
