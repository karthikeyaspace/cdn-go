<!-- https://chatgpt.com/c/672d6f01-c4b0-800e-ba0f-eaa6ec50104a -->
<!-- https://medium.com/geekculture/go-cafe-creating-and-adding-files-to-aws-s3-using-golang-b92eaa5f2081 -->
<!-- https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/ -->

- users can upload index.html and styles.css files to /upload endpoint
- these files are stored in aws s3 bucket
- users can view the static website at /view/{unique-id} endpoint
- the website is served from the s3 bucket