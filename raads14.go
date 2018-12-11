// RAADS14 screens for autism spectrum disorder (ASD).
// https://doi.org/10.1186/2040-2392-4-49
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Answers to questions from the person.
var answers [14]int

// Options for answers to questions.
var options = [4]string{
	"Never true",
	"True only when I was younger than late-adolescence",
	"True only now",
	"True now and when I was younger than late-adolescence",
}

// Questions to ask the person.
var questions = [14]string{
	"It is difficult for me to understand how other people are feeling when we are talking.",
	"Some ordinary textures that do not bother others feel very offensive when they touch my skin.",
	"It is very difficult for me to work and function in groups.",
	"It is difficult to figure out what other people expect of me.",
	"I often don't know how to act in social situations.",
	"I can chat and make small talk with people.",
	"When I feel overwhelmed by my senses, I have to isolate myself to shut them down.",
	"How to make friends and socialize is a mystery to me.",
	"When talking to someone, I have a hard time telling when it is my turn to talk or to listen.",
	"Sometimes I have to cover my ears to block out painful noises (like vacuum cleaners or people talking too much or too loudly).",
	"It can be very hard to read someone's face, hand, and body movements when we are talking.",
	"I focus on details rather than the overall idea.",
	"I take things too literally, so I often miss what people are trying to say.",
	"I get extremely upset when the way I like to do things is suddenly changed.",
}

// Domains for questions.
var domains = [3][]int{
	{12, 0, 8, 3, 10, 11, 13}, // Mentalizing deficits
	{2, 4, 5, 7},              // Social anxiety
	{1, 6, 9},                 // Sensory reactivity
}

// Main program asks questions and displays results.
func main() {
	askall()
	review()
	for !check() {
		review()
	}
	reverse()
	report()
}

// Clear the terminal screen.
func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// Lowercase the first letter in sentences.
func lcfirst(s string) string {
	if s[0:2] == "I " {
		return s
	}
	return strings.ToLower(s[0:1]) + s[1:]
}

// Ask all questions.
func askall() {
	for qn, _ := range questions {
		ask(qn)
	}
}

// Ask and answer the question.
func ask(qn int) {
	var answered bool
	for !answered {
		clear()
		fmt.Printf("[ %d/14 ] \n\n%s\n\n", qn+1, questions[qn])
		for qn, option := range options {
			fmt.Printf("%d: %s\n", qn, option)
		}
		var answer int = -1
		fmt.Println()
		fmt.Scanln(&answer)
		if answer >= 0 && answer <= 3 {
			answers[qn] = answer
			answered = true
		}
	}
}

// Check for validity using the trick question.
func check() bool {
	var prevAnswer int
	for i, answer := range answers {
		// Valid if answers are not all the same
		if i > 0 && answer != prevAnswer {
			return true
		}
		// Notify and allow changing of answers
		if i == 13 {
			fmt.Println()
			fmt.Println("Check your answers and try again")
			fmt.Scanln()
		}
		prevAnswer = answer
	}
	return false
}

// Review answers and change them.
func review() {
	var confirm string
	for confirm != "y" {
		clear()
		for i, question := range questions {
			fmt.Printf("%d. It is [%s] that %s\n", i+1, lcfirst(options[answers[i]]), lcfirst(question))
		}
		fmt.Println()
		fmt.Println("Is this correct?")
		fmt.Println()
		fmt.Println("y: Yes, this is correct.")
		fmt.Println("n: No, I need to change an answer.")
		fmt.Println()
		fmt.Scanln(&confirm)
		if confirm == "n" {
			var qn int = -1
			fmt.Println()
			fmt.Print("What is the number of the question you'd like to change? ")
			fmt.Scanln(&qn)
			if qn >= 1 && qn <= 14 {
				ask(qn - 1)
			}
		}
	}
}

// Reverse the answer to trick question.
func reverse() {
	answers[5] = 3 - answers[5]
}

// Report results.
func report() {
	clear()
	var total int
	var domainTotals [3]int
	for i, domain := range domains {
		for _, n := range domain {
			total += answers[n]
			domainTotals[i] += answers[n]
		}
	}

	var line = "================================================================================"
	var line2 = "--------------------------------------------------------------------------------"

	fmt.Println(line)
	if total >= 24 && total <= 38 {
		fmt.Println("Within close range of those with an autism spectrum disorder diagnosis. (Median: 32)")
	}
	if total >= 9 && total <= 22 {
		fmt.Println("Within close range of those with an attention deficit hyperactivity disorder diagnosis. (Median: 15)")
	}
	if total >= 6 && total <= 17 {
		fmt.Println("Within close range of those with a different psychiatric disorder diagnosis. (Median: 12)")
	}
	if total >= 0 && total <= 6 {
		fmt.Println("Within close range of those without any psychiatric disorder diagnosis. (Median: 3)")
	}
	if total >= 14 {
		fmt.Println("Further investigation of autism spectrum disorder is recommended.")
	}
	fmt.Println(line)

	fmt.Println("Total:", total)
	fmt.Println()
	fmt.Println("Mentalizing deficits:", domainTotals[0])
	fmt.Println("Social anxiety:", domainTotals[1])
	fmt.Println("Sensory reactivity:", domainTotals[2])
	fmt.Println(line2)
	fmt.Scanln()
	fmt.Println()

	fmt.Println(line)
	fmt.Println("Total")
	fmt.Println(line)
	fmt.Printf("You:\t\t\t\t%d\n\n", total)
	fmt.Println("ASD:\t\t\t\t32\t(8-42)")
	fmt.Println("ASD male:\t\t\t30\t(0-42)")
	fmt.Println("ASD female:\t\t\t34\t(9-42)")
	fmt.Println("ADHD:\t\t\t\t15\t(0-42)")
	fmt.Println("ADHD male:\t\t\t15\t(0-36)")
	fmt.Println("ADHD female:\t\t\t15\t(0-42)")
	fmt.Println("Other psychiatric:\t\t11\t(0-39)")
	fmt.Println("Other psychiatric male:\t\t11.5\t(0-33)")
	fmt.Println("Other psychiatric female:\t12\t(0-39)")
	fmt.Println("Non-psychiatric:\t\t3\t(0-29)")
	fmt.Println("Non-psychiatric male:\t\t3\t(0-19)")
	fmt.Println("Non-psychiatric female:\t\t2.5\t(0-29)")
	fmt.Println(line2)
	fmt.Scanln()
	fmt.Println()

	fmt.Println(line)
	fmt.Println("Mentalizing deficits")
	fmt.Println(line)
	fmt.Printf("You:\t\t\t\t%d\n\n", domainTotals[0])
	fmt.Println("ASD male:\t\t\t15\t(0-21)")
	fmt.Println("ASD female:\t\t\t18\t(3-21)")
	fmt.Println("ADHD male:\t\t\t7\t(0-21)")
	fmt.Println("ADHD female:\t\t\t8\t(0-21)")
	fmt.Println("Other psychiatric male:\t\t4\t(0-21)")
	fmt.Println("Other psychiatric female:\t5\t(0-21)")
	fmt.Println("Non-psychiatric male:\t\t1\t(0-13)")
	fmt.Println("Non-psychiatric female:\t\t0\t(0-19)")
	fmt.Println(line2)
	fmt.Scanln()
	fmt.Println()

	fmt.Println(line)
	fmt.Println("Social anxiety")
	fmt.Println(line)
	fmt.Printf("You:\t\t\t\t%d\n\n", domainTotals[1])
	fmt.Println("ASD male:\t\t\t9.7\t(0-12)")
	fmt.Println("ASD female:\t\t\t9\t(0-12)")
	fmt.Println("ADHD male:\t\t\t3\t(0-12)")
	fmt.Println("ADHD female:\t\t\t3\t(0-12)")
	fmt.Println("Other psychiatric male:\t\t4\t(0-11)")
	fmt.Println("Other psychiatric female:\t2\t(0-12)")
	fmt.Println("Non-psychiatric male:\t\t1\t(0-8)")
	fmt.Println("Non-psychiatric female:\t\t0\t(0-10)")
	fmt.Println(line2)
	fmt.Scanln()
	fmt.Println()

	fmt.Println(line)
	fmt.Println("Sensory reactivity")
	fmt.Println(line)
	fmt.Printf("You:\t\t\t\t%d\n\n", domainTotals[2])
	fmt.Println("ASD male:\t\t\t6\t(0-9)")
	fmt.Println("ASD female:\t\t\t8\t(3-9)")
	fmt.Println("ADHD male:\t\t\t3\t(0-9)")
	fmt.Println("ADHD female:\t\t\t3\t(0-9)")
	fmt.Println("Other psychiatric male:\t\t2\t(0-9)")
	fmt.Println("Other psychiatric female:\t3\t(0-9)")
	fmt.Println("Non-psychiatric male:\t\t0\t(0-6)")
	fmt.Println("Non-psychiatric female:\t\t0\t(0-9)")
	fmt.Println(line2)
}
