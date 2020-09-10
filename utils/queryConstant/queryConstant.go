package utils

const (
	INSERT_USER             = "insert into tb_user (user_id,user_firstname,user_lastname,user_address,user_phone,user_image,user_status) values (?,?,?,?,?,?,?)"
	INSERT_AUTH             = "insert into tb_auth(username,password,user_id) values (?,?,?)"
	SELECT_USER_BY_ID       = "select auth.username, auth.password,user.user_image,user.user_poin,user.user_status, user.user_firstname, user.user_lastname, user.user_phone, user.user_address from tb_user user inner join tb_auth auth on auth.user_id = user.user_id where user.user_id = ?"
	SELECT_ALL_USER         = "SELECT tu.user_id,tu.user_firstname,tu.user_lastname,tu.user_address,tu.user_phone,tu.user_poin,tu.user_image,tu.user_status,ta.username,ta.password,ta.user_level_id,ta.user_status FROM tb_user tu JOIN tb_auth ta ON tu.user_id=ta.user_id WHERE tu.user_firstname LIKE ? LIMIT %s ,%s"
	UPDATE_USER             = "UPDATE tb_user SET user_firstname=?,user_lastname=?,user_address=?,user_phone=?,user_image=?,user_status=? WHERE user_id=?"
	UPDATE_AUTH             = "UPDATE tb_auth SET username=?,password=? WHERE user_id=?"
	DELETE_AUTH             = "UPDATE tb_auth SET user_status = NA WHERE user_id = ?"
	LOGIN                   = "select user_id, user_level_id, user_status from tb_auth where username = ? and password= ?;"
	SELECT_AUTH_BY_USERNAME = "SELECT tu.user_id,tu.user_firstname,tu.user_lastname,tu.user_address,tu.user_phone,tu.user_poin,tu.user_email,tu.user_image,tu.user_status,ta.username,ta.password,ta.user_status FROM tb_user tu JOIN tb_auth ta ON tu.user_id=ta.user_id WHERE ta.username = ?"
	INSERT_WALLET           = "INSERT INTO tb_wallet (wallet_id,user_id) values (?,?)"
	SELECT_WALLET_USER_ID   = "SELECT tw.wallet_id,tw.amount,tu.user_id,tu.user_firstname,tu.user_lastname,tu.user_address,tu.user_phone,tu.user_poin FROM tb_wallet tw JOIN tb_user tu ON tw.user_id=tu.user_id WHERE tu.user_id = ?"
	UPDATE_AMOUNT_WALLET    = "UPDATE tb_wallet SET amount = ? WHERE user_id = ?"
	UPDATE_POIN_USER        = "UPDATE tb_user SET user_poin = ? WHERE user_id = ?"
	INSERT_TOP_UP           = "INSERT INTO tb_top_up (top_up_id,top_up_amount,user_id,top_up_date) VALUES (?,?,?,?)"
	UPDATE_STATUS_TOP_UP    = "UPDATE tb_top_up SET top_up_status = ? WHERE user_id = ?"
	//STORE
	INSERT_STORE                = "INSERT INTO tb_store (store_id,store_name,store_category_id,store_address,store_owner,store_username,store_password) VALUES (?,?,?,?,?,?,?)"
	SELECT_STORE_BY_ID          = "SELECT ts.store_id,ts.store_name,ts.store_address,ts.store_owner,ts.store_status,ts.store_username,ts.store_images,tsc.store_category_id,tsc.store_category_name FROM tb_store ts JOIN tb_store_category tsc ON ts.store_category_id=tsc.store_category_id WHERE store_id = ?"
	SELECT_ALL_STORE            = "SELECT ts.store_id,ts.store_name,ts.store_address,ts.store_owner,ts.store_status,ts.store_username,ts.store_password,ts.store_images,tsc.store_category_id,tsc.store_category_name FROM tb_store ts JOIN tb_store_category tsc ON ts.store_category_id=tsc.store_category_id WHERE ts.store_status = 'A'"
	UPDATE_STORE                = "UPDATE tb_store SET store_name=?,store_category_id=?,store_address=?,store_owner=?,store_username=?,store_password=?,store_images=?,qr_path=? WHERE store_id=?"
	DELETE_STORE                = "UPDATE tb_store SET store_status = NA WHERE store_id = ?"
	STORE_AUTH                  = "SELECT * FROM tb_store WHERE store_username = ?"
	INSERT_STORE_CATEGORY       = "INSERT INTO tb_store_category VALUES(?,?)"
	SELECT_STORE_CATEGORY_BY_ID = "SELECT * FROM tb_store_category WHERE store_category_id=?"
	SELECT_ALL_STORE_CATEGORY   = "SELECT * FROM tb_store_category"
	UPDATE_STORE_CATEGORY       = "UPDATE tb_store_category SET store_category_name=? WHERE store_category_id=?"
	DELETE_STORE_CATEGORY       = "UPDATE tb_store_category SET store_category_status = 'NA' WHERE store_category_id = ?"
	//QUERY PRODUCT CATEGORY
	INSERT_PRODUCT_CATEGORY       = "INSERT INTO tb_product_category VALUES (?,?)"
	SELECT_PRODUCT_CATEGORY_BY_ID = "SELECT * FROM tb_product_category WHERE product_category_id = ?"
	SELECT_ALL_PRODUCT_CATEGORY   = "SELECT * FROM tb_product_category"
	UPDATE_PRODUCT_CATEGORY       = "UPDATE tb_product_category SET product_category_name=? WHERE product_category_id=?"
	DELETE_PRODUCT_CATEGORY       = "UPDATE tb_product_category SET product_category_status = 'NA' WHERE product_category_id=? "
	//QUERY PRODUCT
	SELECT_ALL_PRODUCT_BY_STORE = "select pp.product_id, p.product_name, p.product_stock, p.product_status,pc.product_category_name, pp.price, pp.date_modified from tb_product_price pp inner join tb_product p on p.product_id = pp.product_id inner join tb_product_category pc on p.product_category_id = pc.product_category_id inner join (select product_id, max(date_modified) as maxDate from tb_product_price group by product_id) pj on pp.product_id = pj.product_id and pp.date_modified = pj.maxDate where p.store_id= ? AND p.product_status = 'A'"
	INSERT_PRODUCT              = "INSERT INTO tb_product(product_id,store_id,product_name,product_category_id,product_stock) VALUES (?,?,?,?,?)"
	INSERT_PRODUCT_PRICE        = "INSERT INTO tb_product_price VALUES(?,?,?,?)"
	SELECT_PRODUCT_BY_ID        = "select pp.product_id, p.product_name, p.product_stock,p.product_status,pc.product_category_id,pc.product_category_name, pp.price, pp.date_modified from tb_product_price pp inner join tb_product p on p.product_id = pp.product_id inner join tb_product_category pc on p.product_category_id = pc.product_category_id inner join (select product_id, max(date_modified) as maxDate from tb_product_price group by product_id) pj on pp.product_id = pj.product_id and pp.date_modified = pj.maxDate where p.product_id= ? AND p.product_status = 'A'"
	UPDATE_PRODUCT_WITH_PRICE   = "UPDATE tb_product SET product_name = ?,product_stock = ?,product_category_id = ? WHERE product_id = ?"
	DELETE_PRODUCT              = "UPDATE tb_product SET product_status = 'NA' WHERE product_id = ?"

	//FEEDBACK
	GET_ALL_FEEDBACK   = "SELECT * FROM tb_feedback"
	GET_FEEDBACK_BY_ID = "SELECT * FROM tb_feedback WHERE feedback_id = ?"
	POST_FEEDBACK      = "INSERT INTO tb_feedback(feedback_id, store_id, user_id, feedback_value, feedback_created) VALUES (?, ?, ?, ?, ?)"
	UPDATE_FEEDBACK    = "UPDATE tb_feedback SET store_id=?, user_id=?, feedback_value=?, feedback_created=? WHERE feedback_id=?"
	DELETE_FEEDBACK    = "DELETE FROM tb_feedback WHERE feedback_id = ?"

	//POINT
	GET_ALL_POINT     = "SELECT * FROM tb_poin"
	GET_POINT_BY_ID   = "SELECT * FROM tb_poin WHERE poin_id = ?"
	POST_POINT        = "INSERT INTO tb_poin(poin_id, store_id) VALUES(?, ?)"
	UPDATE_POINT      = "UPDATE tb_poin SET store_id=? WHERE product_id=?"
	DELETE_POINT      = "DELETE FROM tb_poin WHERE poin_id = ?"
	UPDATE_USER_POINT = "UPDATE tb_user SET user_poin = user_poin + 1 WHERE user_id=?"

	//RATING
	GET_ALL_RATING   = "SELECT * FROM tb_rating"
	GET_RATING_BY_ID = "SELECT * FROM tb_rating WHERE rating_id = ?"
	POST_RATING      = "INSERT INTO tb_rating(rating_id, store_id, user_id, rating_value, rating_description, rating_created) VALUES (?, ?, ?, ?, ?, ?)"
	UPDATE_RATING    = "UPDATE tb_rating SET store_id=?, user_id=?, rating_value=?, rating_description=?, rating_created=? WHERE rating_id=?"
	DELETE_RATING    = "DELETE FROM tb_rating WHERE rating_id = ?"
)
