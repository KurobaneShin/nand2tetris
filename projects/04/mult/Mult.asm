@2
M=0

@i
M=0

(LOOP)
@i
D=M

@0
D=D-M
@END
D;JGT // if D>=0, jump to END

@i
D=M
@2
M=D+M
@i
M=M+1
@LOOP
0;JMP // jump to LOOP

(END)
@END
0;JMP // loop forever
