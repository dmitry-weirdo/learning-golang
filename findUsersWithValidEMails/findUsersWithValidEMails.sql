-- 1517. Find Users With Valid E-Mails

select
*
from Users u
where u.mail ~ '^[A-Za-z][A-Za-z0-9_.-]*@leetcode[.]com$'

/*
    ^ - string start
    $ - string end
    \. - escape the dot to use it as a character
    * - appears any times
    within [], we can skip the escaping of `_` `.` `-` characters. `-` is the last to NOT work as range
*/
