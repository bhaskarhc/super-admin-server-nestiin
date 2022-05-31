CREATE TABLE IF NOT EXISTS admin_users (id INT NOT NULL,
                                               first_name varchar(250) NOT NULL,
                                                                       last_name varchar(250) NOT NULL,
                                                                                              email varchar(250) NOT NULL unique,
                                                                                                                          password varchar NOT NULL,
                                                                                                                                           mobile varchar(250) NOT NULL,
                                                                                                                                                               created_at TIMESTAMPTZ,
                                                                                                                                                               update_at TIMESTAMPTZ,
                                                                                                                                                               is_verified boolean default false,
                                                                                                                                                                                           PRIMARY KEY (id));


INSERT INTO admin_users (id,first_name,last_name,email,password,mobile)
VALUES(100001,
       'bhaskar',
       'hc',
       'bhaskar@nestiin.com',
       'admin@123',
       '9666334149');