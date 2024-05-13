if [ -n "$1" ]; then
    script1="./$1 $2"
else
    script1=""
fi

if [ -n "$3" ]; then
    script2="./$3 $4"
else
    script2=""
fi

if [ -n "$5" ]; then
    script3="./$5 $6"
else
    script3=""
fi

if [ -n "$7" ]; then
    script4="./$7 $8"
else
    script4=""
fi

if [ -n "$9" ]; then
    script5="./$9 ${10}"
else
    script5=""
fi

if [ -n "${11}" ]; then
    script6="./${11} ${12}"
else
    script6=""
fi

if [ -n "${13}" ]; then
    script7="./${13} ${14}"
else
    script7=""
fi

if [ -n "${15}" ]; then
    script8="./${15} ${16}"
else
    script8=""
fi

if [ -n "${17}" ]; then
    script9="./${17} ${18}"
else
    script9=""
fi

if [ -n "${19}" ]; then
    script10="./${19} ${20}"
else
    script10=""
fi

if [ -n "${21}" ]; then
    script11="./${21} ${22}"
else
    script11=""
fi

if [ -n "${23}" ]; then
    script12="./${23} ${24}"
else
    script12=""
fi

if [ -n "${25}" ]; then
    script13="./${25} ${26}"
else
    script13=""
fi

if [ -n "${27}" ]; then
    script14="./${27} ${28}"
else
    script14=""
fi

if [ -n "${29}" ]; then
    script15="./${29} ${30}"
else
    script15=""
fi

if [ -n "${31}" ]; then
    script16="./${31} ${32}"
else
    script16=""
fi

paste <($script1) <($script2) <($script3) <($script4) <($script5) <($script6) <($script7) <($script8) <($script9) <($script10) <($script11) <($script12) <($script13) <($script14) <($script15) <($script16)
