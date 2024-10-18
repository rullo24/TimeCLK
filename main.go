package main

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Function to update the timer
func update_time_and_label(p_time_label *widget.Label, p_seconds_since_timer_start *uint64) {
	*p_seconds_since_timer_start += 1                                          // incrementing the var each second
	p_time_label.SetText(strconv.FormatUint(*p_seconds_since_timer_start, 10)) // string formatting for base-10 uint64
}

func conv_seconds_to_HH_mm_SS_string(seconds_since_timer_start uint64) {
	return
}

func main() {
	var a fyne.App = app.New()
	var w fyne.Window = a.NewWindow("TimeCLK")

	// Label to show the timer
	var p_time_label *widget.Label = widget.NewLabel("0")
	var current_time_obj time.Time = time.Now() // init current time
	var seconds_since_timer_start uint64 = 0
	var timer_on_bool bool = false

	go func() { // running time update in background (every so often)
		time.Sleep(20 * time.Millisecond) // updating every certain length of time
		current_time_obj = time.Now()
	}()

	// Start button to start the timer
	var p_start_button *widget.Button = widget.NewButton("Start", func() { // start button to start timer
		go func() { // function to run in external thread
			timer_on_bool = true // setting at the beginning of start button press
			for timer_on_bool {  // only continues whilst the timer is active
				time.Sleep(1 * time.Second)                                     // update each second
				update_time_and_label(p_time_label, &seconds_since_timer_start) // updating the timer label on the screen
			}
		}()
	})

	// Reset button to reset the timer
	var p_reset_button *widget.Button = widget.NewButton("Reset", func() {
		timer_on_bool = false         // stopping counting
		seconds_since_timer_start = 0 // resetting the timer
		p_time_label.SetText("0")     // resetting the text
	})

	// Layout for the app
	var content *fyne.Container = container.NewVBox(p_time_label, p_start_button, p_reset_button)
	w.SetContent(content)
	w.ShowAndRun() // starting the window loop
}
