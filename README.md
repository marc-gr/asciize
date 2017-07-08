# Asciize

  <a href="https://github.com/marc-gr/asciize/releases/latest"><img src="https://img.shields.io/github/release/marc-gr/asciize.svg?style=flat-square"/></a>
  <a href="https://godoc.org/github.com/marc-gr/asciize"><img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square"/></a>
  <a href="https://travis-ci.org/marc-gr/asciize"><img src="https://img.shields.io/travis/marc-gr/asciize.svg?style=flat-square"/></a>
  <a href="https://goreportcard.com/report/github.com/marc-gr/asciize"><img src="https://goreportcard.com/badge/github.com/marc-gr/asciize?style=flat-square&x=1"/></a>
  <a href="https://github.com/marc-gr/asciize/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square"/></a>
  
  This is a port to pure go of the original ruby gem [asciiart](https://github.com/nodanaonlyzuul/asciiart)

## Installation

If you want to install the binary, just go to the [releases](https://github.com/marc-gr/asciize/releases) page and download the latest available for your system.

If you have go installed:

To install only the cli:

```bash
$ go get -u github.com/marc-gr/asciize/cmd/asciize
```

To also install the code:
```bash
$ go get -u github.com/marc-gr/asciize
```

## Usage

### In Code

```go
f, _ := os.Open("./sample.png")
img, err := png.Decode(f)
if err != nil {
	panic(err)
}
a := asciize.NewAsciizer(asciize.Colored(true))
s, err := a.Asciize(img)
if err != nil {
	panic(err)
}
fmt.Println(s)
```

For more go to [the examples folder](examples)

### In The Command Line

Local Files

```bash
$ asciize -src "examples/sample.png"
```
```
                                           oooooooooooooooo                                          
                                 oo%%%###########################%%oo                                
                           oo%#########################################%oo                           
                       o%###################################%%ooooo%%######%o     oooooo             
         oo%%%%oo   o%####%%oo%%%%%ooo%#################%oo%%MWWWWWM%%oo%#####%oo#########o          
      o%########%oo####%oo%MWWWWWWWWWW%%o%############oo%WWWWWWWWWWWWWWW%o%#####%o%#%#######o        
     %#####%###oo####%o%WWWWWWWWWWWWWWWWW%o%########%oMWWWWWWWWWWWWWWWWWWW%o######o   o######o       
    %####o    o%####o%WWWWWWWWWWWWWWWWWWWWWo%######%oWWWWWWWWWWWWWWWWWWWWWW%o######%   o######       
   o####o    o#####%oWWWM%%MWWWWWWWWWWWWWWWWo######oWWWoo  oMWWWWWWWWWWWWWWWo#######% o#######       
    ####%o  o######oWWM      MWWWWWWWWWWWWWWo%####%oWW       MWWWWWWWWWWWWWWo%#######%o######o       
    o#####%o#######oWWo    o oWWWWWWWWWWWWWW%o####%oWW    oM %WWWWWWWWWWWWWWo%########%o###%o        
      %###o%#######oMW%    %o%WWWWWWWWWWWWWWo%#####oMW%     %WWWWWWWWWWWWWWMo##########ooo           
         oo########%oWWMooooMWWWWWWWWWWWWWW%o######%oMWWMMMWWWWWWWWWWWWWWWMo%##########%             
          ##########%oWWWWWWWWWWWWWWWWWWWW%o####%%%##o%WWWWWWWWWWWWWWWWWMoo#############o            
         o###########%o%WWWWWWWWWWWWWWWM%o%#%o       o%oo%MWWWWWWWWWWM%oo###############%            
         %#############%oo%%MWWWWWWM%%oo%##%           ##%%ooo%%%%ooo%###################            
         ##################%%oooooo%%###%oo%          %%o%###############################o           
         ##############################%o%%%%%%%%%%%%%%%%%o%#############################o           
        o##############################o%%%%%%%%%%%%%%%%%%%o%############################%           
        o##############################oo%%%%%%%%o%%%%%%%%%o##############################           
        o###############################%ooo%%MW%oWWM%oooo%###############################           
        o#################################o%WWWWo%WWWWo###################################           
        o#################################%oWWWWo%WWWWo###################################o          
         ##################################ooMM%oo%WM%o###################################o          
         ###################################%%%%##%oo%####################################o          
         %################################################################################o          
         %################################################################################o          
         o################################################################################o          
         o################################################################################o          
          ################################################################################o          
          ################################################################################o          
          ################################################################################o          
          %###############################################################################o          
          %###############################################################################o o        
   o%%%%%o%###############################################################################%o%%%%%o   
 %%%%%%%%o%###############################################################################%o%%%%%%%  
 oo%%%%%% %################################################################################ %%%%%oo  
  %%%o    %################################################################################    ooo   
          #################################################################################o         
          #################################################################################o         
         o#################################################################################o         
         o#################################################################################o         
         %#################################################################################o         
         %#################################################################################%         
         ##################################################################################%         
         ##################################################################################%         
         ##################################################################################o         
         ##################################################################################o         
         ##################################################################################o         
         ##################################################################################          
         ##################################################################################          
         %################################################################################o          
         o################################################################################           
         o###############################################################################%           
          %##############################################################################            
          o#############################################################################o            
           %###########################################################################%             
            %#########################################################################%              
             %#######################################################################%               
              o#####################################################################o                
                %#################################################################%                  
                 o%#############################################################%o                   
                 o%oo########################################################%oo%%%o                 
              o%%%%%%ooo%#################################################%oo%%%%%%%%o               
             %%%%%%%%%%%o oo%########################################%%o     %%%%%%%%%o              
            %%o%%%%%%%o         oo%%%#########################%%ooo            %%%%%o%%              
            oo%%%%%%                      oooooooooooooo                         %%%%                
              o%o                                                                                    
                                                                                                     
```
   
Remote Images
```bash
 $ asciize -src "http://www.google.com/images/srpr/logo3w.png"
```

Output it as HTML

```bash
$ asciize -c -f html -src "examples/sample.png" > ascii-as-html.html
```
    
_or smaller_

```bash
$ asciize -w 50 -c -f html -src "examples/sample.png" > ascii-as-html.html
```

Get Help

```bash
Usage of ./asciize:
  -c	If set to true the output will be colored (ANSI or HTML)
  -cs string
    	Used to define a custom charset (default " .~:+=o*x^%#@$MW")
  -f string
    	Output format. Can be "text" or "html" (default "text")
  -i	If set to true the charset will be reversed. Can improve results for some images
  -src string
    	Define the source of the image. Can be a local file or a URL
  -w uint
    	Target width (default 100)
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
