

GO
-- list order and items by item serial
-- i refers to item , oi refers to orderItem
CREATE PROCEDURE [dbo].TablesUnpauseByImei (@Imei VARCHAR(100))
AS
BEGIN
	DECLARE @ComputerName VARCHAR(50)
	SELECT @ComputerName = ComName FROM ComUse WHERE Imei = @Imei
	UPDATE dbo.Tables SET pause = 0 , ComputerName = CASE WHEN State = 'Free' THEN NULL ELSE ComputerName END   WHERE ComputerName = @ComputerName
END


GO
ALTER PROCEDURE [dbo].[TablesListByGroupNo](@GroupTableNo int)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;
	DECLARE @TotalCash FLOAT
	SELECT  Tables.Serial ,  Tables.TableNo , TableName , "pause" , "State" ,
	 	ISNULL(StkTr03.PrintTimes , 0) PrintTimes , 
	 	ISNULL(DocDate , '') DocDate , ISNULL(StkTr03.DocNo , '') DocNo ,
		ISNULL(StkTr03.OrderNo , 0) OrderNo ,
		ISNULL(StkTr03.BonNo , 0) BonNo ,
		ISNULL(StkTr03.CustSNo , 0) Guests ,
	  	ISNULL(StkTr03.Serial , 0) HeadSerail ,  ISNULL(StkTr03.WaiterCode ,0) WaiterCode 
		  , ISNULL(StkTr03.AccountSerial ,0) AccountSerial ,
	   	ISNULL((SELECT SUM(Qnt * Price) FROM  StkTr04 WHERE HeadSerial = StkTr03.Serial) ,0) Subtotal ,
		ISNULL(DiscountPercent , 0) DiscountPercent ,
        ISNULL(Tables.ComputerName , '') ComputerName
	FROM  "Tables" 
		LEFT JOIN StkTr03 
		ON StkTr03.TableSerial = Tables.Serial AND ISNULL(TotalCash ,0) = 0 
	
	WHERE Tables.GroupTableNo = @GroupTableNo
END
