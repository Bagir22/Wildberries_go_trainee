package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	ps "github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Shell struct {
	in       io.Reader
	out      io.Writer
	currDir  string
	pipes    bool
	pipeArgs []string
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Can't get current directory")
	}
	shell := &Shell{os.Stdin, os.Stdout, dir, false, []string{}}
	fmt.Fprintln(os.Stdout, "Run shell")
	shell.Run()
	fmt.Fprintln(os.Stdout, "Shell exit")
}

func (s *Shell) Run() {
	scanner := bufio.NewScanner(s.in)
	for {
		currDir := fmt.Sprintf("%s ~ %% ", s.currDir)
		fmt.Print(currDir)

		if scanner.Scan() {
			input := scanner.Text()

			//Завершаем при exit
			if input == "exit" {
				return
			}

			//Обработка команд
			s.ParseCommands(input)
		} else {
			fmt.Println(scanner.Err())
			return
		}
	}
}

func (s *Shell) ParseCommands(input string) {
	//Разделение введенной строки на команды, если используется конвейер
	commands := strings.Split(input, "|")
	for _, cmd := range commands {
		cmd = strings.TrimSpace(cmd)
		parts := strings.Fields(cmd)
		if len(parts) == 0 {
			continue
		}

		//Проверка наличия конвейера
		if len(commands) > 1 {
			s.pipes = true
			s.pipeArgs = append(s.pipeArgs, parts...)
			continue
		}

		//Обработка команд
		switch parts[0] {
		case "cd":
			s.executeCD(parts)
		case "pwd":
			s.executePWD()
		case "echo":
			s.executeEcho(parts)
		case "kill":
			s.kill(parts[1])
		case "exec":
			s.ex(parts[1:])
		default:
			fmt.Println("unknown command")
		}
	}
}

func (s *Shell) executeCD(parts []string) {
	if len(parts) < 2 {
		fmt.Fprintln(s.out, "Invalid cd command") //Если недостаточно аргументов
		return
	}
	err := os.Chdir(parts[1]) //Изменяем директорию
	if err != nil {
		fmt.Fprintln(s.out, "Error: ", err)
	}
	s.currDir, err = os.Getwd()
	if err != nil {
		fmt.Println("Can't get current directory")
	}
}

func (s *Shell) executePWD() {
	dir, err := os.Getwd() //Получаем текущую директорию
	if err != nil {
		fmt.Fprintln(s.out, "Error:", err)
	} else {
		fmt.Fprintln(s.out, dir)
	}
}

func (s *Shell) executeEcho(parts []string) {
	fmt.Fprintln(s.out, strings.Join(parts[1:], " ")) //Выводим echo
}

func (s *Shell) ExecuteCommand(parts []string) {
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdin = s.in
	cmd.Stdout = s.out
	cmd.Stderr = s.out

	if s.pipes {
		cmd.Args = append(s.pipeArgs, cmd.Args...)
	}

	fmt.Println(cmd.Args, parts[1:])
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(s.out, "Error:", err)
	}
	fmt.Println(err)
	//Сбрасываем флаги конвейера
	s.pipes = false
	s.pipeArgs = nil
}

func (s *Shell) kill(arg string) {
	if pid, err := strconv.Atoi(arg); err != nil {
		fmt.Println(arg)
		fmt.Println(err)
	} else if proc, err := ps.FindProcess(pid); err != nil {
		//ошибка в поиске процесса
		fmt.Println(err)
	} else if proc == nil {
		//процесса с таким pid нет
		fmt.Println("kill:", pid, ": процесс не существует")
	} else if proc, err := os.FindProcess(pid); err != nil {
		//ошибка в поиске процесса
		fmt.Println(err)
	} else {
		//Убиваем процесс
		proc.Kill()
	}

}

func (s *Shell) ex(args []string) {

	if len(args) < 1 {
		fmt.Println("err")
	}

	argv := []string{}
	if len(args) > 1 {
		argv = args[1:]
	}

	//вызов exec для args[0]
	if err := syscall.Exec(args[0], argv, os.Environ()); err != nil {
		fmt.Println(err)
	}
}
