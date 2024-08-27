Certainly! Here's an updated version of the README that highlights the usefulness of AppTimer for managing screen time, curbing addiction, and enhancing focus during work:

---

# AppTimer - A Productivity CLI Application

**AppTimer** is a lightweight command-line tool designed to help you manage your time effectively by controlling how long you use specific applications. Whether you're aiming to limit distractions, reduce screen time, or enhance your focus, AppTimer offers a simple yet powerful solution.

## Why AppTimer?

In today's digital age, it's easy to lose track of time while using various applications, whether it's streaming services, social media, or even work-related software. AppTimer helps you take back control by setting strict time limits on your app usage. This can be particularly beneficial if:

- **You're battling screen addiction:** AppTimer can assist in reducing excessive screen time by automatically closing apps after a set period.
- **You want to focus on deep work:** By limiting the time spent on distracting applications, AppTimer helps you maintain focus on important tasks.
- **You're managing your work-life balance:** Set boundaries on how long work-related apps are open, ensuring you don't overextend your work hours.

## Features

- **Launch any application:** Start any application on your system directly from the command line.
- **Set a time limit:** Define how long you want the application to run.
- **Automatic termination:** Once the time limit is reached, the application is automatically closed.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/varrow1/apptimer.git
   ```

2. **Navigate to the project directory:**

   ```bash
   cd apptimer
   ```

3. **Build the application:**

   ```bash
   go build -o apptimer
   ```

4. **Run the application:**

   ```bash
   ./apptimer
   ```

## Usage

1. **Run AppTimer from the command line:**

   ```bash
   ./apptimer
   ```

2. **Input the application you want to launch:**
   
   You will be prompted to enter the name or path of the application you want to launch. For example, to open Netflix, type:

   ```
   netflix
   ```

3. **Set the duration:**
   
   You will then be asked how long you want the application to run. Enter the time in minutes. For example, to run the application for 30 minutes, type:

   ```
   30
   ```

4. **Stay in control:**
   
   AppTimer will launch the application and automatically close it after the specified time. This helps you stay focused on your goals and avoid getting sucked into time-consuming activities.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

This version emphasizes the practical benefits of AppTimer for improving focus and managing screen time, making it clear how the tool can be used to address common productivity challenges.
