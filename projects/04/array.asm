// pointers are variables that holds memory addresses like arr and i, are called pointers
//arr = 100
@100
D=A
@arr
M=D //puts 100 on arr variable

//n=10
@10
D=A
@n
M=D // puts 10 on variable n to iterate the for loop

// initialize i = 0
@i
M=0

(LOOP)

//if (i==n) goto END
@i
D=M
@n
D=D-M
@END
D;JEQ

// RAM[arr+i] = -1
@arr
D=M //gets the value saved in arr variable
@i
A=D+M 
M=-1 //the address affected is the aritmetic result of D+M

//i++
@i
M=M+1

@LOOP
0;JMP

(END)
@END
0;JMP
