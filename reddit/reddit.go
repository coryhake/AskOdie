package main

import (      
        "fmt"  
        "log"
        "code.google.com/p/odie"
        "github.com/jzelinskie/reddit"
)


func frontPageHandler(w *odie.ResponseWriter, req *odie.Request, vars odie.Context) {
        //get headlines
        headlines, err := reddit.DefaultHeadlines()
          if err != nil {
            log.Println(err)
          }
        
        //print to ResponseWriter
        fmt.Fprintf(w, "<h3>Reddit Front Page</h3>") //print
        fmt.Fprintf(w, "<img src=\"http://is.gd/g6YfA1\" alt=\"subreddit header image\"></img>") //print

        for key := range headlines {
          //log.Println("headline title: ", headlines[key].Title) //print
          fmt.Fprintf(w, "%s: <a href=\"%s\">%s</a>", headlines[key].Subreddit, headlines[key].URL, headlines[key].Title) //print
        }
}

func subHandler(w *odie.ResponseWriter, req *odie.Request, vars odie.Context) {
        //get Subreddit
        sub := vars.Get("subreddit")
        subreddit, err := reddit.AboutSubreddit(sub)
          if err != nil {
            log.Println("ERROR IN CALL - AboutSubreddit: ", err)
            fmt.Fprintf(w, "Oops! It looks like that subreddit doesn't exist. Please try again using a valid subreddit.") //print
            return
          }
        
        //get SubredditHeadlines
        headlines, err := reddit.SubredditHeadlines(sub)
          if err != nil {
            log.Println("ERROR IN CALL - SubredditHeadlines: ", err)
            fmt.Fprintf(w, "Oops! It looks like that subreddit doesn't exist. Please try again using a valid subreddit.") //print
            return
          }
        
        //NSFW Content Check
        if subreddit.IsNSFW {
          fmt.Fprintf(w, "Sorry, can't display NSFW content, but you're welcome to follow this direct link.") //print
          fmt.Fprintf(w, "<h3>Subreddit: <a href=\"http://www.reddit.com/%s\">%s</a></h3>", subreddit.URL, subreddit.Name) //print
          return
        }
        
        //print to ResponseWriter
        fmt.Fprintf(w, "<h3>Subreddit: <a href=\"http://www.reddit.com/%s\">%s</a></h3>", subreddit.URL, subreddit.Name) //print
        fmt.Fprintf(w, "<img src=\"%s\" alt=\"subreddit header image\"></img>", subreddit.HeaderImg) //print
        for key := range headlines {
          //log.Println("headline title: ", headlines[key].Title) //print
          fmt.Fprintf(w, "%s: <a href=\"%s\">%s</a>", headlines[key].Subreddit, headlines[key].URL, headlines[key].Title) //print
        }
  
}

func redditorHandler(w *odie.ResponseWriter, req *odie.Request, vars odie.Context) {
        //get Redditor
        user := vars.Get("username")
        redditor, err := reddit.AboutRedditor(user)
          if err != nil {
            log.Println("ERROR IN CALL - AboutRedditor: ", err)
            fmt.Fprintf(w, "Oops! It looks like that redditor doesn't exist. Please try again using a valid username.") //print
            return
          }
        
        //print to ResponseWriter
        fmt.Fprintf(w, "<h3>Details for Redditor: %s</h3>", user) //print
        fmt.Fprintf(w, "ID: %s", redditor.Id) //print
        fmt.Fprintf(w, "Link Karma: %d", redditor.LinkKarma) //print
        fmt.Fprintf(w, "Comment Karma: %d", redditor.CommentKarma) //print
          if redditor.Gold {
              fmt.Fprintf(w, "This user has Gold") //print
          }
          if redditor.Mod {
              fmt.Fprintf(w, "This user is Mod") //print
          }
}


//print handler functionalities to ReponseWriter
func helpHandler(w *odie.ResponseWriter, req *odie.Request, vars odie.Context) {
        
        fmt.Fprintf(w, "<h3>Using the Reddit Plugin</h3>") //print
        fmt.Fprintf(w, "<b>red front:</b><br> No variables needed here. Displays the default Reddit front page. <br>Returned list of links includes the name of the subreddit the comment is from, along with the title and a direct link.<br>Default front page has NSFW content disabled.") //print
        fmt.Fprintf(w, "<b>red sub $subreddit:</b> <br>Enter \"red sub subreddit\", i.e. \"red sub funny\". Displays a list of headlines for that subreddit sorted by popularity.<br>")
        fmt.Fprintf(w, "<b>red user $username:</b> <br>Enter \"red user username\", i.e. \"red user Unidan\". Displays details for the selected redditor, including their unique ID and link/comment karma. Also checks for Gold and Mod status.<br>")
        fmt.Fprintf(w, "<b>but how do it did that?</b><br> Take a look at the code for the <a href=\"http://is.gd/p8bJ9U\">Ask Reddit Pack.</a>") //print
        fmt.Fprintf(w, "<img id=\"gopher\" src=\"http://is.gd/KjLJ03\" alt=\"gopher\" style=\"margin-left:80px;\"></img>") //print
}

//initialize handlers with odie package
func init () {

        odie.Handle("reddit front", frontPageHandler) 
        odie.Handle("reddit sub $subreddit", subHandler) 
        odie.Handle("reddit user $username", redditorHandler)
        odie.Handle("reddit help", helpHandler)
}

func main () {

        odie.SubscribeAndServe(&odie.AppInfo{Name:"Ask Reddit Pack", Author:"Cory Hake"})
}