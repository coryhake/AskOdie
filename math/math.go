package main

import (
        "strconv"
        "fmt"
        "log"
        "math"
        "code.google.com/p/odie"
)

func additionHandler(w *odie.ResponseWriter, req *odie.Request, vars odie.Context) {
        //get variables from input
        var1 := vars.Get("addend1")
        var2 := vars.Get("addend2")
          if var1 == "" {
              return
           }
          if var2 == "" {
           return
          }
        
        //parse strings to int
        addend1, err1 := strconv.Atoi(var1) //parse var1
          if err1 != nil {
            log.Println("ParseInt error1: ", err1)
          }
        addend2, err2 := strconv.Atoi(var2) //parse var2
          if err2 != nil {
            log.Println("ParseInt error2: ", err2)
          }
        
        //sum and print
        sum := addend1 + addend2 //add 
        fmt.Fprintf(w, "%d + %d = %d", addend1, addend2, sum) //print     
}

func subtractionHandler(w *odie.ResponseWriter, req *odie.Request, vars odie.Context) {
        //get variables from input
        var1 := vars.Get("minuend")
        var2 := vars.Get("subtrahend")
          if var1 == "" {
              return
           }
          if var2 == "" {
           return
          }
        
        //parse strings to int
        minuend, err1 := strconv.Atoi(var1)
          if err1 != nil {
            log.Println("ParseInt error1: ", err1)
          }
        subtrahend, err2 := strconv.Atoi(var2)
          if err2 != nil {
            log.Println("ParseInt error2: ", err2)
          }
        
        //subtract and print
        difference := minuend - subtrahend //subtract
        fmt.Fprintf(w, "%d - %d = %d", minuend, subtrahend, difference) //print
}

func multiplicationHandler(w *odie.ResponseWriter, req *odie.Request, vars odie.Context) {
        //get variables from input
        var1 := vars.Get("multiplicand")
        var2 := vars.Get("multiplier")
          if var1 == "" {
              return
           }
          if var2 == "" {
           return
          }
        
        //parse strings to int
        multiplicand, err1 := strconv.Atoi(var1)
          if err1 != nil {
            log.Println("ParseInt error1: ", err1)
          }
        multiplier, err2 := strconv.Atoi(var2)
          if err2 != nil {
            log.Println("ParseInt error2: ", err2)
          }
        
        //multiply and print
        product := multiplicand * multiplier
        fmt.Fprintf(w, "%d x %d = %d", multiplicand, multiplier, product) //print
}

func divisionHandler(w *odie.ResponseWriter, req *odie.Request, vars odie.Context) {
        //get variables from input
        var1 := vars.Get("dividend")
        var2 := vars.Get("divisor")
          if var1 == "" {
              return
           }
          if var2 == "" {
           return
          }
        
        //parse strings to int
        dividend, err1 := strconv.ParseFloat(var1, 64)
          if err1 != nil {
            log.Println("ParseFloat error1: ", err1)
          }
        divisor, err2 := strconv.ParseFloat(var2, 64)
          if err2 != nil {
            log.Println("ParseFloat error2: ", err2)
          }
        
        //divide and print
        quotient := dividend / divisor
        fmt.Fprintf(w, "%.0f / %.0f = %.2f", dividend, divisor, quotient) //print
}

func exponentHandler(w *odie.ResponseWriter, req *odie.Request, vars odie.Context) {
        //get variables from input
        var1 := vars.Get("base")
        var2 := vars.Get("exponent")
          if var1 == "" {
              return
          }
          if var2 == "" {
           return
          }
        
        //parse strings to int
        base, err1 := strconv.ParseFloat(var1, 64)
          if err1 != nil {
            log.Println("ParseFloat error1: ", err1)
          }
        exponent, err2 := strconv.ParseFloat(var2, 64)
          if err2 != nil {
            log.Println("ParseFloat error2: ", err2)
          }
        
        //calculate and print
        solution := math.Pow(base, exponent)
        fmt.Fprintf(w, "%.0f ^ %.0f = %.0f", base, exponent, solution) //print
}

//print handler functionalities to ReponseWriter
func helpHandler(w *odie.ResponseWriter, req *odie.Request, vars odie.Context) {

        fmt.Fprintf(w, "<b>add:</b><br> Enter \"add $addend1 $addend2\" i.e. \"add 2 3\". <br>Math Pack will sum the addends and display your result") //print
        fmt.Fprintf(w, "<b>subtract:</b><br> Enter \"subtract $minuend $subtrahend\" i.e. \"subtract 3 2\". <br>Math Pack will find the difference between minuend and subtrahend and display your result") //print
        fmt.Fprintf(w, "<b>multiply:</b><br> Enter \"multiply $multiplicand $multiplier\" i.e. \"multiply 2 3\". <br>Math Pack will find the product of multiplicand and multiplier and display your result") //print
        fmt.Fprintf(w, "<b>divide:</b><br> Enter \"divide $dividend $divisor\" i.e. \"divide 6 3\". <br>Math Pack will find the quotient of dividend and divisor and display your result") //print
        fmt.Fprintf(w, "<b>exponent:</b><br> Enter \"exponent $base $exponent\" i.e. \"exponent 2 3\". <br>Math Pack will find the product of base to the power of exponent and display your result") //print
        fmt.Fprintf(w, "<b>but how do it maths?</b><br> Take a look at the code for the <a href=\"http://is.gd/XaWMzG\">Math Pack.</a>") //print
        fmt.Fprintf(w, "<img id=\"gopher\" src=\"http://is.gd/KjLJ03\" alt=\"gopher\" style=\"margin-left:70px;\"></img>") //print
}

//initialize handlers with odie package
func init () {
  
        odie.Handle("add $addend1 $addend2", additionHandler) 
        odie.Handle("subtract $minuend $subtrahend", subtractionHandler)
        odie.Handle("multiply $multiplicand $multiplier", multiplicationHandler)
        odie.Handle("divide $dividend $divisor", divisionHandler)
        odie.Handle("exponent $base $exponent", exponentHandler)
        odie.Handle("math help", helpHandler)
}

func main () {

        odie.SubscribeAndServe(&odie.AppInfo{Name:"Math Pack", Author:"Cory Hake"})
}