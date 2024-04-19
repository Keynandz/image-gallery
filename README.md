# KeynandZ Image Gallery (RESPONSIVE)

![IMG](https://i.ibb.co/sPP8hWz/image-2024-04-18-102232265.png)

KeynandZ Image Gallery is a simple web application that allows users to view and upload images to a gallery. This application utilizes Go technology for the backend server, Minio as the file storage, and HTML/CSS/JavaScript for the frontend view.

## Key Features

- View the uploaded image gallery.
- Upload new images to the gallery.

## Prerequisites

Before running this project, make sure you have the following prerequisites:

- Go (version 1.16 or newer) installed on your computer.
- Minio server installed and configured.
- Go module enabled.

## How to Use

1. Clone this repository to your computer:

    ```
    git clone https://github.com/Keynandz/image-gallery.git
    ```

2. Navigate to the project directory:

    ```
    cd keynandz-image-gallery
    ```

3. Create a `.env` file and set up the configuration as needed. Example content of `.env`:

    ```
    MINIO_ENDPOINT=http://localhost:9000
    MINIO_ACCESS_KEY=minio_access_key
    MINIO_SECRET_KEY=minio_secret_key
    MINIO_BUCKET=image-gallery
    ```

4. Install Go dependencies using Go module:

    ```
    go mod tidy
    ```

5. Run the application:

    ```
    go run main.go
    ```

6. Open a browser and access `http://localhost:9000` to view the image gallery.

