# Program 4 - Concurrent Image Downloader
## Requirements
* Program should reads a list of image URLs from a file.
* Two versions of the downloader implemented:
    -Sequential Version: Downloads and saves each image one after the other.
    -Concurrent Version: Downloads and saves images concurrently using goroutines.
* Downloaded images are saved to disk with unique names - end of image url
* Proper error handling for failed downloads.
* Time taken by each version to complete the downloads are printed.
* Largely written by ChatGPT
* **Sequential Time: 1.6829861s**
* **Concurrent Time: 168.1735ms**

|   #   | File                                                                          | Description                                                                   |
| :---: | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
|   1   | [main.go](main.go)                                                            | [main](main.go)                                                               |
|   2   | [Image 1](image_1701962140001104800.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640) | [Image 1](image_1701962140001104800.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640) |
|   3   | [Image 2](image_17019621400577903000)                                         | [Image 2](image_17019621400577903000)                                         |
|   4   | [Image 3](image_1701962140244218600)                                          | [Image 3](image_1701962140244218600)                                          |
|   5   | [Image 4](image_1701962140244218600.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640) | [Image 4](image_1701962140244218600.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640) |

* Image 1
![Image 1](image_1701962140001104800.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640)
* Image 2
![Image 2](image_17019621400577903000)
* Image 3
![Image 3](image_1701962140244218600)
* Image 4
![Image 4](image_1701962140244218600.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640)
