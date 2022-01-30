use hoge;

#users
UPDATE hoge.users SET email=replace(email,left(email,instr(email,'@')-1),CONCAT('dummy_email_',id));
UPDATE hoge.users SET pass='dummypassword';
UPDATE hoge.users SET address=CONCAT('dummy_address_',id);
UPDATE hoge.users SET company=CONCAT('dummy_company',id);
UPDATE hoge.users SET name=CONCAT('dummy_name_',id);
UPDATE hoge.users SET tell=CONCAT('dummy_tell_',id);