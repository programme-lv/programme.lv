tmp_path = mktempdir()
gen_dir = joinpath("test-source","subtask-gen")

test_dir = joinpath("testing-dst","tests")
mkpath(test_dir)

Threads.@threads for gen_filename in readdir(gen_dir)
    generator = splitext(gen_filename)[1]
    gen_exe = joinpath(tmp_path, generator)
    gen_src = joinpath(gen_dir, gen_filename)
    
    println("compiling $gen_src")
    run(`g++ -o $gen_exe $gen_src`)

    println("generating 10 tests using $gen_exe")
    for i = 1:10
        test_num = string(i,base=10,pad=3)
        test_name = "$(generator)_$(test_num).in"
        test_path = joinpath(test_dir,test_name)
        open(f->write(f, read(`$gen_exe $i`, String)), test_path, "w")
    end
end

sol_exe = joinpath(tmp_path, "solution")
sol_src = joinpath("test-source","solution.cpp")

run(`g++ -o $sol_exe $sol_src`)

Threads.@threads for test_filename in readdir(test_dir)
    test_path = joinpath(test_dir, test_filename)
    test_name = splitext(test_filename)[1]
    ans_path = joinpath(test_dir, "$test_name.ans")
    println("generating $ans_path")
    open(f->write(f, read(pipeline(`cat $test_path`,`$sol_exe`),String)), ans_path,"w")
end

chk_exe = joinpath("testing-dst", "checker")
chk_src = joinpath("test-source","checker.cpp")

run(`g++ -o $chk_exe $chk_src`)