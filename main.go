package main

import (
	"fmt"
	"os"

	"github.com/santedev/gothBase/template"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices       [][]string
	textInput     textinput.Model
	outputMessage string
	spinner       spinner.Model
	steps         int
	length        int
	limit         int
	cursor        int
	selected      map[int]string
	message       string
}

const (
	templateTemplStr =
		"                                               \n" +
		" (               )    )   (                   \n" +
		" )\\ )         ( /( ( /( ( )\\     )        (   \n" +
		"(()/(     (   )\\()))\\()))((_) ( /(  (    ))\\  \n" +
		" /(_))_   )\\ (_))/((_)/((_)_  )(_)) )\\  /((_) \n" +
		"(_)) __| ((_)| |_ | |(_)| _ )((_)_ ((_)(_))   \n" +
		"  | (_ |/ _ \\|  _|| ' \\ | _ \\/ _` |(_-</ -_)  \n" +
		"   \\___|\\___/ \\__||_||_||___/\\__,_|/__\\/___|  \n"
)

var (
	grayStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	redStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("210"))
	greenStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("120"))
	boldGrayStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Bold(true)
	boldStyle     = lipgloss.NewStyle().Bold(true)
	basicCmd      = fmt.Sprintf(`%s%s%s%s%s%s%s%s%s%s%s%s%s%s`,
		moveRandL,
		grayStyle.Render(`". Move "`),
		boldGrayStyle.Render("up"), grayStyle.Render("/"), boldGrayStyle.Render("down"),
		grayStyle.Render(`" | "`),
		boldGrayStyle.Render("k"), grayStyle.Render("/"), boldGrayStyle.Render("j"),
		grayStyle.Render(`". Select option "`),
		boldGrayStyle.Render("Enter"),
		grayStyle.Render(`" | "`),
		boldGrayStyle.Render("Space"),
		grayStyle.Render(`"`),
	)
	moveRandL = fmt.Sprintf("%s%s%s%s%s%s%s%s",
		grayStyle.Render(`Next "`),
		boldGrayStyle.Render("right"),
		grayStyle.Render(`" | "`),
		boldGrayStyle.Render("n"),
		grayStyle.Render(`". Back "`),
		boldGrayStyle.Render("left"),
		grayStyle.Render(`" | "`),
		boldGrayStyle.Render("b"))
	tempOptions = make([]string, 0, 6)
	errMsg      = ""
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.steps {
		case 1:
			return stepCommand1(m, msg)
		case 2:
			return stepCommand2(m, msg)
		case 3:
			return stepCommand3(m, msg)
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tea.Msg:
		if msg == "done" {
			m.outputMessage = greenStyle.Render("Project created successfully!")
			return m, tea.Quit
		}
		if msg == "err" {
			m.outputMessage = errMsg
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	switch m.steps {
	case 1:
		return stepView1(m)
	case 2:
		return stepView2(m)
	case 3:
		return stepView3(m)
	default:
		return ""
	}
}

func stepView1(m model) string {
	s := templateTemplStr + "\n\n"

	if m.message != "" {
		s += "\n" + m.message + "\n"
	}
	s += m.textInput.View()
	s += "\n\n" + fmt.Sprintf(`%s%s%s%s%s%s%s`,
		grayStyle.Render(`Press "`),
		boldGrayStyle.Render("ctrl+c"),
		grayStyle.Render(`" | "`),
		boldGrayStyle.Render("q"),
		grayStyle.Render(`" to quit. Press "`),
		boldGrayStyle.Render("Enter"),
		grayStyle.Render(`" to continue.`))
	s += boldGrayStyle.Render(fmt.Sprintf(" %d/%d", m.steps, m.limit))
	return s
}

func stepView2(m model) string {
	// The header
	s := "\n" + boldStyle.Render("Select the options that your project will have:") + "\n"

	// Iterate over our choices
	for i, choice := range m.choices[m.steps-1] {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		checked := " " // not selected
		if _, ok := m.selected[i+m.length]; ok {
			checked = "x" // selected!
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// Add the message if available
	if m.message != "" {
		s += "\n" + m.message + "\n"
	}

	s += "\n" + basicCmd + grayStyle.Render(".")
	s += boldGrayStyle.Render(fmt.Sprintf(" %d/%d", m.steps, m.limit))
	return s
}

func stepView3(m model) string {
	s := "\n" + `do you want to create your project with name`
	name := ""
	if m.textInput.Value() == "" {
		m.textInput.SetValue("project")
	}
	if m.textInput.Value() != "" {
		name = boldStyle.Render(m.textInput.Value())
	}
	s += ` "` + boldStyle.Render(name) + `"` + "\n\n"
	s += " with option/s:\n"
	if len(m.selected) == 0 {
		m.selected[0] = m.choices[1][0]
	}
	if len(m.selected) > 0 && m.outputMessage == "" {
		for _, value := range m.selected {
			s += "  -" + boldStyle.Render(value) + "\n"
		}
	}
	if m.outputMessage != "" && len(tempOptions) == 0 {
		for _, value := range m.selected {
			tempOptions = append(tempOptions, value)
		}
	}
	if m.outputMessage != "" && len(tempOptions) > 0 {
		for _, value := range tempOptions {
			s += "  -" + boldStyle.Render(value) + "\n"
		}
	}

	s += "\n" + grayStyle.Render(`Press key "`) + boldGrayStyle.Render("Y/n ") + grayStyle.Render(`"to continue or cancel.`)
	s += " " + moveRandL + grayStyle.Render(`".`)
	s += boldGrayStyle.Render(fmt.Sprintf(" %d/%d", m.steps, m.limit))
	if m.steps == m.limit && m.outputMessage != "" {
		s += "\n\n" + m.outputMessage + "\n\n"
		s += "Do not exit the CLI yet " + m.spinner.View()
	}
	return s
}

func stepCommand1(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	m, cmd := defaultCmdInput(m, msg.String())
	if cmd != nil {
		return m, cmd
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func stepCommand2(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	m, cmd := defaultCmd(m, msg.String())
	if cmd != nil {
		return m, cmd
	}
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}

	case "down", "j":
		if m.cursor < len(m.choices[m.steps-1])-1 {
			m.cursor++
		}

	case "enter", " ":
		_, ok := m.selected[m.length+m.cursor]
		if ok {
			delete(m.selected, m.length+m.cursor)
		} else {
			if ok := m.handleSelections(
				m.choices[m.steps-1][m.cursor],
				m.steps-1,
				"default",
				"minimalistic"); ok {
				m.selected[m.length+m.cursor] = m.choices[m.steps-1][m.cursor]
			}
		}
	}
	return m, nil
}

func stepCommand3(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	m, cmd := defaultCmd(m, msg.String())
	if cmd != nil {
		return m, cmd
	}
	switch msg.String() {
	case "y", "Y":
		tempOptions = make([]string, 0, 6)
		if m.textInput.Value() == "" {
			m.textInput.SetValue("project")
		}
		m.outputMessage = "Your project started to be build"
		return m, tea.Batch(m.spinner.Tick, m.createTemplate)

	case "n", "N":
		tempOptions = make([]string, 0, 6)
		m.outputMessage = redStyle.Render("Canceled creation of the project")
		return m, tea.Quit
	}

	return m, nil
}

func defaultCmd(m model, msg string) (model, tea.Cmd) {
	switch msg {
	case "ctrl+c", "q":
		return m, tea.Quit

	case "n", "right":
		if m.steps < m.limit {
			m.length += len(m.choices[m.steps-1])
			m.steps++
			m.cursor = 0
		}

	case "b", "left":
		if m.steps > 1 {
			if m.length-len(m.choices[m.steps-2]) >= 0 {
				m.length -= len(m.choices[m.steps-2])
			}
			m.steps--
			m.cursor = 0
		}

	}
	return m, nil
}

func defaultCmdInput(m model, msg string) (model, tea.Cmd) {
	switch msg {
	case "ctrl+c":
		return m, tea.Quit

	case "enter":
		if m.steps < m.limit {
			m.steps++
		}

	case "ctrl+u":
		if m.steps > 1 {
			if m.length-len(m.choices[m.steps-2]) >= 0 {
				m.length -= len(m.choices[m.steps-2])
			}
			m.steps--
			m.cursor = 0
		}

	}
	return m, nil
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink)
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Project name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	sp := spinner.New()

	return model{
		choices: [][]string{
			{},
			{"default", "minimalistic", "auth", "database sql", "Dockerfile", "javascript", "htmx", "tailwind"},
			{},
		},
		limit:     3,
		steps:     1,
		length:    0,
		textInput: ti,
		spinner:   sp,

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the choices slice, above.
		selected: make(map[int]string),
	}
}

func (m model) handleSelections(current string, step int, standalonArr ...string) bool {
	standalon := ""
	for _, std := range standalonArr {
		if current == std {
			standalon = current
			break
		}
	}
	if standalon == "" {
	loop:
		for _, value := range m.selected {
			for _, std := range standalonArr {
				ok, _ := contains(m.choices[step], value)
				if value == std && ok {
					standalon = value
					break loop
				}
			}
		}
	}

	if standalon == current {
		for _, value := range m.selected {
			ok, key := contains(m.choices[step], value)
			if value != standalon && ok {
				delete(m.selected, m.length+key)
			}
		}
		return true
	}

	for _, value := range m.selected {
		if len(standalon) > 0 && standalon != value {
			ok, key := contains(m.choices[step], value)
			if ok {
				delete(m.selected, m.length+key)
			}
		}
	}
	return standalon == ""
}

func contains(slice []string, str string) (bool, int) {
	for key, value := range slice {
		if value == str {
			return true, key
		}
	}
	return false, -1
}

func (m *model) createTemplate() tea.Msg {
	sl := make(map[string]bool, len(m.selected))
	for _, value := range m.selected {
		sl[value] = true
	}

	projectName := m.textInput.Value()
	selections := template.Selections{
		ProjectName:  projectName,
		Default:      sl["default"],
		Minimalistic: sl["minimalistic"],
		Auth:         sl["auth"],
		DatabaseSQL:  sl["database sql"],
		Javascript:   sl["javascript"],
		Dockerfile:   sl["Dockerfile"],
		Htmx:         sl["htmx"],
		Tailwind:     sl["tailwind"],
	}
	err := template.GenerateTemplate(selections)
	if err != nil {
		errMsg = redStyle.Render(err.Error())
		return tea.Msg("err")
	}
	return tea.Msg("done")
}
