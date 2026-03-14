# 🌈 Kid World - Premium PWA 
**"Making Learning Fun and Ad-Free"**

Kid World एक हल्का और सुरक्षित एजुकेशनल वेब ऐप है, जिसे मैंने खास तौर पर छोटे बच्चों के लिए बनाया है। यह ऐप बच्चों को खेल-खेल में हिंदी वर्णमाला, अंग्रेजी Alphabets और Numbers सिखाने में मदद करता है।

---

### 🌟 Key Features
- **PWA Enabled:** यह एक असली ऐप की तरह फोन में इंस्टॉल हो सकता है और ऑफलाइन भी काम करता है।
- **Ad-Free Experience:** बच्चों के लिए पूरी तरह सुरक्षित, कोई भी फालतू विज्ञापन नहीं।
- **Multi-Language Support:** हिंदी (अ-ज्ञ) और अंग्रेजी (A-Z) का बेहतरीन उच्चारण (Native Speech Logic)।
- **Admin Insights:** डेटा-ड्रिवन अप्रोच के लिए कस्टम एडमिन डैशबोर्ड (Tracking correct/wrong clicks)।

---

### 🚀 Deployment Workflow (CLI Driven)
इस प्रोजेक्ट को मैंने पूरी तरह से कमांड लाइन (CLI) के जरिए डिप्लॉय किया है:

1. **Repo Setup:**
   - `git init` (लोकल फोल्डर को गिट से जोड़ा)
   - `gh repo create Kid-World --public --source=. --remote=origin` (GitHub CLI से रिपोजिटरी बनाई)

2. **Commit & Push:**
   - `git add .`
   - `git commit -m "Initial launch of Kid World Premium PWA"`
   - `git push -u origin main`

3. **Live Hosting (GitHub Pages):**
   - `gh repo edit --enable-pages`
   - `gh repo edit --pages-branch main`

---

### 🛠️ Tech Stack
- **Frontend:** HTML5, CSS3, JavaScript (PWA Service Workers)
- **Backend:** Go (Golang) - *Admin Tracking API*
- **Database:** SQLite - *User Engagement Data*
- **Deployment:** GitHub Pages & GitHub CLI

---

### 📊 Admin Monitoring
मैंने इस प्रोजेक्ट में **DevOps Principles** का इस्तेमाल करते हुए एक एडमिन डैशबोर्ड जोड़ा है, जो `Go` बैकएंड से जुड़ा है। यह हमें यह समझने में मदद करता है कि बच्चे ऐप के साथ कैसे इंटरैक्ट कर रहे हैं।

---
**Developed by Himanshu** *Transitioning from Telecom to DevOps with a focus on Automation & Product Excellence.*
