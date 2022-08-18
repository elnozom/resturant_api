
GO
ALTER PROC [dbo].[StkMs01InsertUpdate]
(
@ItemCode INT = NULL,
@GroupCode INT = NULL ,
@BarCode VARCHAR(20)= NULL,
@Name VARCHAR(200)= NULL,
@MinorPerMajor INT= NULL,
@AccountSerial int= NULL,
@ActiveItem bit= NULL,
@ItemTypeID int = NULL,
@ItemHaveSerial bit= NULL,
@MasterItem bit= NULL,
@StoreCode int = NULL,
@LastBuyPrice real= NULL,
@POSTP real= NULL,
@POSPP real = NULL,
@Ratio1  real= NULL,
@Ratio2 real = NULL,
@Disc1 real= NULL,
@Disc2 real= NULL,
@PriceBefore real= NULLl
)
AS
BEGIN
	DECLARE @ItemAccountSerial int;
	declare @ItemSerial int  
	declare @InsertItemAccountQuery VARCHAR(300)
	SET @ItemSerial = (SELECT Serial FROM StkMs01 WHERE ItemCode = @ItemCode AND GroupCode = @GroupCode)
	--SET @InsertItemAccountQuery = 'INSERT INTO ItemAccount (AccountSerial , ItemSerial , Disc1 , Disc2 , LastBuyPrice , FinalPrice) VALUES (' + @AccountSerial + ' , ' + @ItemSerial + ','  + @Disc1 + ',' + @Disc2 + ',' + @PriceBefore + ' ,' + @POSTP + ')'
	SET @InsertItemAccountQuery = CONCAT('INSERT INTO ItemAccount (AccountSerial , ItemSerial , Disc1 , Disc2 , LastBuyPrice , FinalPrice) VALUES (' ,@AccountSerial , ' , ' ,  @ItemSerial ,  ',' , @Disc1  ,',' ,  @Disc2 ,',' , @PriceBefore , ' ,'  , @LastBuyPrice , ')')
	-- check if the item is exist to update it nad return so if the next condition is true
	-- the rest of the code is not executed
	IF @ItemSerial IS NOT NULL
		BEGIN
		  UPDATE StkMs01 
			SET BarCode  =@BarCode ,
			 POSName  =@Name ,
			 ItemName  =@Name ,
			 ItemTypeID =@ItemTypeID ,
			 MinorPerMajor =@MinorPerMajor ,
			 AccountSerial =@AccountSerial ,
			 ActiveItem =@ActiveItem ,
			 ItemHaveSerial =@ItemHaveSerial ,
			 MasterItem =@MasterItem 

		WHERE 
		ItemCode = @ItemCode AND GroupCode = @GroupCode




		UPDATE StkMs02
					SET ItemSerial = @ItemSerial,
					StoreCode = @StoreCode,
					LastBuyPrice = @LastBuyPrice,
					AvrPrice = @LastBuyPrice,
					POSTP = @POSTP,
					POSPP = @POSPP,
					Ratio1 = @Ratio1,
					Ratio2 = @Ratio2,
					Percnt1 = 1,
					Percnt2 =1 

		WHERE 
		ItemSerial = @ItemSerial

		--check if the record exists into item account to update it or insert it
		SET @ItemAccountSerial = (SELECT Serial FROM ItemAccount WHERE AccountSerial = @AccountSerial AND ItemSerial = @ItemSerial)
		IF @ItemAccountSerial IS NOT NULL
			BEGIN
				UPDATE ItemAccount SET AccountSerial = @AccountSerial , ItemSerial = @ItemSerial , Disc1 = @Disc1 , Disc2 = @Disc2 , LastBuyPrice =  @LastBuyPrice , FinalPrice = @POSTP WHERE Serial = @ItemAccountSerial
			END
		ELSE 
			BEGIN
				EXEC (@InsertItemAccountQuery)
			END

        SELECT @ItemSerial serial
		RETURN 
	END
-- if this block executed that means its a new product 
INSERT INTO StkMs01 (
    ItemCode , GroupCode , BarCode , POSName , ItemName ,ItemTypeID,MinorPerMajor,AccountSerial,
	ActiveItem,ItemHaveSerial,MasterItem
)
 VALUES 
(
    @ItemCode , @GroupCode , @BarCode , @Name , @Name ,@ItemTypeID,@MinorPerMajor,
	@AccountSerial,@ActiveItem ,@ItemHaveSerial,@MasterItem
)

set @ItemSerial = SCOPE_IDENTITY()   


INSERT INTO StkMs02
             (ItemSerial, StoreCode,  LastBuyPrice, AvrPrice,   POSTP, POSPP, Ratio1, Ratio2, Percnt1, Percnt2)
VALUES   (@ItemSerial,@StoreCode,@LastBuyPrice,@LastBuyPrice,@POSTP,@POSPP,@Ratio1,@Ratio2,1,1)

--// insert item accont
EXEC (@InsertItemAccountQuery)




SELECT @ItemSerial serial
END





DROP PROC IF EXISTS ItemsListBo

GO
CREATE  PROCEDURE ItemsListBo(
	@name VARCHAR(200),
	@groupCode INT,
	@priceFrom REAL,
	@priceTo REAL,
	@dateFrom VARCHAR(200),
	@dateTo VARCHAR(200)
)
AS
BEGIN
    SELECT i.ItemName , i.ImagePath , i.estimated_time, g.GroupName , id.POSPP , i.created_at
        FROM StkMs01 i
        JOIN GroupCode g ON i.GroupCode = g.GroupCode
        JOIN StkMs02 id ON i.Serial = id.ItemSerial
		WHERE 
			i.ItemName LIKE CONCAT('%' , dbo.ISEMPTY(@name ,i.ItemName) , '%')
			AND g.GroupCode = dbo.ISZERO(@groupCode , g.GroupCode)
			AND id.POSPP >= dbo.ISZERO(@priceFrom , id.POSPP)
			AND id.POSPP <= dbo.ISZERO(@priceTo , id.POSPP)

END






DROP PROC IF EXISTS ItemsEditAdd
GO
CREATE  PROCEDURE ItemsEditAdd(
	@serial INT,
	@name VARCHAR(200),
	@nameEn VARCHAR(200),
	@groupCode INT,
	@bardCode VARCHAR(200),
	@price real,
	@estimatedTime INT
)
AS
BEGIN
    IF @serial = 0
	BEGIN
		INSERT INTO StkMs01 (
			ItemName,
			ItemNameEn,
			GroupCode , 
			BarCode ,
			estimated_time
		) VALUES (
			@name,
			@nameEn,
			@groupCode , 
			@bardCode ,
			@estimatedTime
		)

		INSERT INTO StkMs02(
			ItemSerial,
			POSPP
		) VALUES (
			SCOPE_IDENTITY(),
			@price
		)
	END
	ELSE
	BEGIN
		UPDATE StkMs01 SET 
			ItemNamE =  @name,
			ItemNameEN =  @nameEn,
			GroupCode  =  @groupCode , 
			BarCode =  @bardCode ,
			estimated_time =  @estimatedTime
		WHERE Serial = @serial
		UPDATE StkMs02 SET POSPP = @price WHERE ItemSerial = @serial
	END

	SELECT SCOPE_IDENTITY() serial

END

