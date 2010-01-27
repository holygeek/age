#!/bin/sh
. ./test-lib.sh

seconds_per_year=$((365 * 24 * 60 * 60))
seconds_per_week=$((7 * 24 * 60 * 60))
seconds_per_day=$((24 * 60 * 60))

ok_equals '1 year duration'  "`age -d $seconds_per_year`"  '1 year'
ok_equals '0 seconds since now' "$(age -s $(date +%s))" '0 seconds'
ok_equals '1 year 1 second duration'  "`age -d $(($seconds_per_year + 1))`"  '1 year 1 second'
ok_equals '1 year 1 week' "`age -d $(($seconds_per_year + $seconds_per_week))`" '1 year 7 days'
ok_equals '2 days' "`age -d $(($seconds_per_day * 2))`" '2 days'

file="$work_dir/test"
touch $file
ok "Age of given file" age -f $file

exit $exit_value
