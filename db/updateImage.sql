CREATE PROC UpdateImage @mainGroup INT , @subGroup INT  ,@product INT , @image VARCHAR(500)
AS 
BEGIN 
    IF @product = 0
        UPDATE GroupCode SET ImagePath = @image WHERE GroupCode = @subGroup
    ELSE
        UPDATE StkMs01 SET ImagePath = @image WHERE ItemCode = @product
        
END

