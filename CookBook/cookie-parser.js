// ใช้ตัวแปรเป้น const เพราะเราไม่ต้องการจะเปลี่ยนค่าตัวแปรสำคัญๆพวกนี้กลางทางแน่นอน เพราะมันคือตัว config หลักๆเลยนั่นเอง !
const express = require("express"); // เสมือน import package มาใช้งาน
const app = express(); // สร้าง Express Application ลองกด ctrl + คลิกเข้าไปดูในไส้ใน
const cookieParser = require("cookie-parser");

// MiddleWare คือตัวกลางที่ใช้ในการจัดการข้อมูลก่อนที่จะผ่านเข้ามายัง request 
// เราลง middleware cookieParser() เอาไว้สำหรับอ่าน header cookie ไม่อย่างนั้นมันจะหาไม่เจอและพังตลอดนั่นเอง
app.use(cookieParser())

// path สมมติที่เราตั้งขึ้นมาเองซึ่งไม่ได้เป็น path จริงๆตามโครงสร้าง folder แต่มันคือ path
// ที่เราใช้เรียกบน url เว็บไซต์นั่นเอง

//parameter req คือ header ของ client ที่ส่งเข้ามา ซึ่งอาจจะดู cookies, session ที่เข้ามาก็ได้
//parameter res คือ สิ่งที่ server Node เราจะตอบสนองข้อมูลกลับไปนั่นเอง
app.get("/", (req, res) => {
  // server จะสามารถส่งทั้ง header ต่างๆหรือจะตัวหนังสือ json อะไรก็ได้กลับไป
  res.send("Hello World");
});

app.get("/setCookies", (req, res) => {
  // res.cookies.[ชื่อcookie] ก็คือระบุว่าเมื่อมีการเข้ามาในเว็บ server จะ response ส่งcookies ไปให้นั่นเอง
  // เวลาหน่วยเป็น millisecond
  res.cookie("myName", "LinjingYun12@Clarrissa", {
    maxAge: 1000 * 60 * 60
  });
  // เมื่อเราต้องการให้ API ตัดการเชื่อมต่อควรใส่ res.end() เอาไว้เพื่อตัดการเชื่อมทันทีไม่อย่างนั้นมันจะหน้าขาวหมุนติ้วๆไปจนกว่า timeout
  console.log(req.cookies["myName"]);
  res.end();
});

// server จะรันที่ port 3000 หรือ port ใดๆก็ตามใจเราและ callback จะทำงานเมื่อ
app.listen(3000, () => {
  console.log("Server Listen At 3000");
});