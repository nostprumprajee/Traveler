// route สำหรับเช็คการ login โดย  :ตัวแปร นั้นหมายถึงชื่อตัวแปรที่ใช้ในการรับค่านั้นเองซึ่งเรียกว่า param !
app.get("/login/:loginData/:food", (req, res) => {
    console.log('Login Check!')
    console.log(req.params.loginData)
    console.log(req.param.food)
  });
  
  // route สำหรับ query string โดย path เริ่มต้นจะเป็น domain:port/api/path/target?key=value&key=value&...
  // ตัวเริ่มตนหลังจาก path ของ route และปิดท้ายpath ด้วย ? ตามด้วยค่า key ของตัวแปรตัวแรก = value และตัวแปรถัดๆไปคั่นด้วยเครื่องหมาย &
  app.get("/login", (req, res) => {
    console.log('Login Query String Check!')
    console.log(req.query) // ได้ query string ทั้งหมด
    console.log(req.query.userName) // เลือกเฉพาะ key
    console.log(req.query.password)
  });