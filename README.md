# 🚀 ASCII Art Web Project

## 🌟 Overview
This project aims to create a Go program that generates graphical representations of strings using ASCII characters. It features multiple banner styles and the ability to write output to files.

## 🔥 Features
- **Input Handling**: Accepts strings with letters, numbers, spaces, and special characters.
- **Banner Styles**: Supports various ASCII banner styles for customization.
- **File Output**: Option to save the ASCII art to a specified file using command-line flags.
- **Best Practices**: Follows good coding practices and includes comprehensive unit tests.

## 🛠 Installation
To run this project, ensure you have Go installed. Download it from [the official Go website](https://golang.org/dl/).

Clone the repository and navigate to the project directory.

```bas
git clone https://github.com/AniMar0/ascii-web-oop.git
```
```bas
cd ascii-art-output
```


## 💻 Usage
Run the program with the following command structure:
```bas
go run . [OPTION] [STRING] [BANNER]
```

## 🧪 Examples
1. Basic usage with string input:
```bas
go run . "Hello"
```
2. Output to a file:
```bas
go run . --output=banner.txt "Hello" standard
```
3. Specifying a banner style:
```bas
go run . --output=banner.txt "Hello There!" shadow
```
4. Using special characters:
```bas
go run . --output=banner.txt $\$ shadow
```
5. Including newlines in input:
```bas
go run . "Hello\nThere"
```
> ⚠️ in this case if you press `enter` on your keyboard it is considered as '\n' and you can't use "\n" (if you do it will be printed as it is)

## ⚙️ Command-Line Options
- `--output=<fileName.txt>`: Specify the output file for the ASCII art.
- `[STRING]`: The string to be converted into ASCII art.
- `[BANNER]`: (Optional) The style of the ASCII banner (e.g., standard, shadow, etc.).

## 📜 Usage Message
If the command format is incorrect, the program displays:
```bas
Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --output=<fileName.txt> something standard
```

## 💻 Tech Stack

**🦄 Languages**:   Go (Golang), Markdown<br> 

**💼 Tools**:   VSCode, Git<br> 

**📦 Packages**:  Only the standard Go packages are allowed <br>

## 📜 License
This project is licensed under the MIT License.

## 🌟 Contributing
Contributions are welcome! Feel free to submit issues or pull requests for improvements or bug fixes.

## 📝 Acknowledgements
Special thanks to my peers especially @melayachi and @bdouirat for their insights and reviews. Thanks to the Go community for providing excellent documentation and support.
