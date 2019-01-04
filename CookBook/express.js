var express = require("express"); // เสมือน import package มาใช้งาน
var app = express(); // สร้าง Express Application ลองกด ctrl + คลิกเข้าไปดูในไส้ใน

// path สมมติที่เราตั้งขึ้นมาเองซึ่งไม่ได้เป็น path จริงๆตามโครงสร้าง folder แต่มันคือ path
// ที่เราใช้เรียกบน url เว็บไซต์นั่นเอง

//parameter req คือ header ของ client ที่ส่งเข้ามา ซึ่งอาจจะดู cookies, session ที่เข้ามาก็ได้
//parameter res คือ สิ่งที่ server Node เราจะตอบสนองข้อมูลกลับไปนั่นเอง
app.get("/", (req, res) => {
  // server จะสามารถส่งทั้ง header ต่างๆหรือจะตัวหนังสือ json อะไรก็ได้กลับไป
  res.send("Hello World");
});

// server จะรันที่ port 3000 หรือ port ใดๆก็ตามใจเราและ callback จะทำงานเมื่อ
app.listen(3000, () => {
    console.log('Server Listen At 3000')
});