clear
echo Compiling: Virtual Machine
g++ svm.cpp -o compiled/svm.out


echo Running: $1.scpb
echo =============================
printf "\n"
printf "\n"

compiled/svm.out ../Testing/scpb/$1.scpb

printf "\n"
printf "\n"
echo =============================
echo Sanitizing: svm.out
rm compiled/svm.out
printf "\n"
