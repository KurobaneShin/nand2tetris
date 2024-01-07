// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen
// by writing 'black' in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen by writing
// 'white' in every pixel;
// the screen should remain fully clear as long as no key is pressed.

(LOOP)
@KEYBOARD
D=M // D = keyboard input
@CHANGESCREEN
D;JEQ // if no key pressed set screen to white = 0
D=-1 // key pressed set screen to black

(CHANGESCREEN)
@PARAM
M=D // save the param
@color // -1 OR 0
D=D-M
@LOOP
D;JEQ // if PARAM = color stop

@PARAM
D=M
@color
M=D // color = PARAM

@SCREEN
D=A // D = SCREEN address
 @8192
 D=D+A // D=Byte just past last screen address
 @i
 M=D // i=SCREEN address

(COLORSCREEN)
@i
D=M-1
M=D
@LOOP
D;JLT // check if i<0 if so go to LOOP

@color
D=M
@i
A=M
M=D // set screen addres to color
@COLORSCREEN
0;JMP // repeat COLORSCREEN





