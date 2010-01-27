export exit_value=0
say_ok() {
    echo "OK $*"
}

say_not_ok() {
    echo "Not OK $*"
}

test_run() {
    #echo  "  CHECK:  $1"
    eval  2>/dev/null 1>/dev/null "$1"
    eval_ret="$?"
    return 0
}

report() {
    exit_status=$1; shift
    name="$*"
    if [ "$exit_status" = 0 -a "$eval_ret" = 0 ]; then
	say_ok "$name"
    else
	say_not_ok "$name"
	exit_value=$(($exit_value + 1))
    fi
}

ok_equals() {
    name="$1"
    left="$2"
    right="$3"

    test_run "test '$left' = '$right'"
    report $? $name
}

ok() {
    name=$1
    shift
    test_run "$*"
    report $? $name
}
