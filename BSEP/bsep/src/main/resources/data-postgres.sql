insert into authority(id, name) values (nextval('my_seq_authority'), 'ROLE_ADMIN');
insert into authority(id, name) values (nextval('my_seq_authority'), 'ROLE_USER');

insert into users(id, user_email,password,first_name,last_name,phone_number,type_of_user) values (nextval('my_seq_users'),'marko.jaksic@gmail.com','$2y$12$UnNvKAnFOA2dVKomUIXtN.yMHYaBuibaB157JScCP/5.n.xgkvf5K','Marko','Jaksic','065-322-3211',1);