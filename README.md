
# 🎨 Ascii Art Web

## 🌟 Project Overview

Welcome to **Ascii Art**, where text meets creativity! 🖌️ This Go-powered command-line tool turns your boring text into eye-popping, stylized ASCII art. Whether you're looking to spice up your terminal output or create some retro-inspired text art, Ascii Art has got you covered! 🌈 Now, with added **web-stylization** and **Dockerization**, it’s easier than ever to showcase your masterpieces to the world!

## ✨ Features

- **🎨 Custom Font Magic**: Your text comes to life with a custom font defined in `standard.txt`. Each letter and symbol has its unique personality, just waiting to jump onto your screen!
- **📝 Multiline Marvel**: Want to create an entire masterpiece? Use `\\n` to craft stunning multiline ASCII art!
- **💻 Command-Line Wizardry**: Just type in your text and let the magic happen. Perfect for adding flair to your scripts or surprising your friends!
- **🌐 Web-Ready Styling**: Your ASCII art isn’t just for terminals anymore! We’ve made it web-friendly with a stylized web interface, so you can share your creations in any browser.
- **🐳 Dockerized Simplicity**: We've Dockerized the project using multi-stage builds. Run your ASCII art tool anywhere with Docker, ensuring it’s easy to deploy and maintain!

## 🚀 How It Works

1. **🔍 Font Loading**: We dive into the `standard.txt` file to fetch the artistic blueprints for your text. Each character gets its own 8-line high makeover.
   
2. **🧩 Input Processing**: Simply input your text at the command line or through the web interface. Want a new line? Just sprinkle a `\\n` in there, and voilà!

3. **🎭 ASCII Art Rendering**: We assemble your characters line by line, creating a dazzling display of ASCII artistry right in your terminal or browser.

## 🎉 Usage

### 🛠️ Command-Line Fun

To unleash the magic, type:

```bash
go run main.go "YourTextHere"
```

- Replace `"YourTextHere"` with the text you want to dazzle the world with.
- Use `\\n` to create new lines in your masterpiece.

### 🌐 Web Stylization

Now you can access and style your ASCII art creations via a web interface! To start the web version:

1. **Run the Dockerized version** using:
   ```bash
   docker run -p 8080:8080 your_docker_image
   ```
2. Visit `http://localhost:8080` in your browser, where you can input text and see your stylish ASCII art in a beautiful web display.

### 🌈 Example

```bash
go run main.go "Hello\\nWorld"
```

This command will light up your screen with:

```
✨ASCII art representation of "Hello"✨
✨ASCII art representation of "World"✨
```

Or, for the web version:

Visit `http://localhost:8080` after running the Docker container and enter "Hello\nWorld" in the input box to see the magic happen!

## 🖼️ Font File (`standard.txt`)

The magic happens inside `standard.txt`, where each character's stylish ASCII look is carefully crafted across 9 lines. We calculate where each character’s art begins and ends, ensuring your text is always dressed to impress. ✨

## 🐳 Dockerizing with Multi-Stage Builds

This project is **Dockerized** for easy deployment. We use a multi-stage Docker build process to ensure a small and efficient image:

- **Stage 1**: We use a Go environment to build the application from source.
- **Stage 2**: We use a minimal Alpine Linux environment to run the compiled application, serving the web interface on port 8080.

This ensures that the final image is lightweight and only contains what’s necessary to run the ASCII art application in both terminal and web modes.

### To build and run the Docker image:

**you can simply use ./docker.sh**
```bash
./docker.sh
```

## ⚠️ Error Handling

- If `standard.txt` decides to play hide and seek, we’ll let you know with a friendly error message and safely exit stage left. 🎭
- Got empty input? No worries! We’ll just print a nice, clean blank line for you.

## 👥 Contributors

This project wouldn’t be possible without the amazing work of our contributors:

- **zdiouri**
- **srm**
- **aayyada**

## 🤝 Contributions

Got ideas to make Ascii Art even cooler? Fork this project and share your awesomeness with a pull request! 🚀

## 📜 License

Feel free to remix, share, and enjoy! This project is licensed under the [MIT License](LICENSE). 🎉
