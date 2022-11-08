# clean previous build
rm -fr build
rm -fr result

# create build, result dir
mkdir build
mkdir result

# compile generator, solution, checker
(
    echo "g++ source/checker.cpp -o result/checker"; 
    echo "g++ source/solution.cpp -o build/solution";
    echo "g++ source/generator.cpp -o build/generator";
) | parallel -t

# create tests
mkdir result/tests
(cd ./result/tests && ../../build/generator)

ls result/tests | parallel -t "build/solution < result/tests/{} > result/tests/{}.a"