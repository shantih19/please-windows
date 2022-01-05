subinclude("//test/build_defs")

please_repo_e2e_test(
    name = "hidden_excluded_dir_test",
    repo = "test_repo",
    plz_command = "plz query alltargets //... > out.txt",
    expected_output = {"out.txt": "//:foo"},
)