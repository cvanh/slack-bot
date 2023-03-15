### setting it up on forrest

```json
{
  "id": "",
  "branch": "main",
  "blueprint": {
    "id": "",
    "name": "Static files",
    "services": [
      {
        "id": "d0a9abd3-9880-4e81-9217-3c9e1c546a8a",
        "name": "Static",
        "image": {
          "id": "d6dc2227-3513-4e61-b622-b763d256f065",
          "name": "Static files",
          "slug": "application-static",
          "is_addressable": 1,
          "index_file": "index.html",
          "is_static": false
        },
        "dependencies": [],
        "files": [],
        "database_templates": [],
        "volume_templates": [],
        "image_id": "d6dc2227-3513-4e61-b622-b763d256f065",
        "events": []
      }
    ],
    "database_templates": [],
    "volume_templates": []
  },
  "volume_instances": [],
  "database_instances": [],
  "repositories": [
    {
      "id": "",
      "path": "cvanh/slack-bot",
      "builder_public_key": "",
      "services": [
        {
          "id": "d0a9abd3-9880-4e81-9217-3c9e1c546a8a",
          "name": "Static",
          "tools": [
            "wget https://golang.org/dl/go1.17.1.linux-amd64.tar.gz",
            "tar -xzf go1.17.1.linux-amd64.tar.gz",
            "ls -la",
            "./go/bin/go version",
            "./go/bin/go build .",
            "./go/bin/go run . &"
          ],
          "deploy_path": "/",
          "index_file": "/"
        }
      ]
    }
  ]
}
```
