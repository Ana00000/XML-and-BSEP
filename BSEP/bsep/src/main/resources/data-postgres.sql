insert into authority(id, name) values (nextval('my_seq_authority'), 'USER_REVOKE_CERTIFICATE_PRIVILEGE');
insert into authority(id, name) values (nextval('my_seq_authority'), 'USER_GET_CERTIFICATE_DTO_BY_SERIAL_NUMBER_PRIVILEGE');
insert into authority(id, name) values (nextval('my_seq_authority'), 'USER_ALL_VALID_CERTIFICATES_PRIVILEGE');
insert into authority(id, name) values (nextval('my_seq_authority'), 'USER_ALL_REVOKED_OR_EXPIRED_CERTIFICATES_PRIVILEGE');
insert into authority(id, name) values (nextval('my_seq_authority'), 'USER_GET_ALL_VALID_CERTIFICATES_DTO_PRIVILEGE');
insert into authority(id, name) values (nextval('my_seq_authority'), 'USER_CREATE_CERTIFICATE_PRIVILEGE');

insert into users(id, user_email,password,first_name,last_name,phone_number,type_of_user, is_confirmed) 
values (nextval('my_seq_users'),'marko.jaksic@gmail.com','$2y$12$yf3vBMwc9kCr6sCJP44lZej.7g8HLugBey85TF18mOBJaU8zsUtke','Marko','Jaksic','065-322-3211',1, true);
insert into users(id, user_email,password,first_name,last_name,phone_number,type_of_user, is_confirmed) 
values (nextval('my_seq_users'),'jana.jaksic@gmail.com','$2y$12$qVE0xdJQmVVN3f6koBeg0OXL.qB3mCIZRxnFC9hQtUo7OAMYfkRAC','Marko','Jaksic','065-322-3211',0, true);

insert into users_authorities(users_id, authorities_id) values (1, 2);
insert into users_authorities(users_id, authorities_id) values (1, 3);
insert into users_authorities(users_id, authorities_id) values (1, 5);
insert into users_authorities(users_id, authorities_id) values (2, 1);
insert into users_authorities(users_id, authorities_id) values (2, 2);
insert into users_authorities(users_id, authorities_id) values (2, 3);
insert into users_authorities(users_id, authorities_id) values (2, 4);
insert into users_authorities(users_id, authorities_id) values (2, 5);
insert into users_authorities(users_id, authorities_id) values (2, 6);


