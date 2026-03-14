<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kid World - Welcome!</title>
    <style>
        body {
            font-family: 'Comic Sans MS', 'Chalkboard SE', sans-serif;
            background-color: #87CEEB;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            background-image: radial-gradient(circle, #fff 10%, transparent 11%), radial-gradient(circle, #fff 10%, transparent 11%);
            background-size: 60px 60px;
            background-position: 0 0, 30px 30px;
        }
        .login-card {
            background-color: white;
            padding: 40px;
            border-radius: 25px;
            box-shadow: 0 10px 25px rgba(0,0,0,0.2);
            text-align: center;
            max-width: 400px;
            width: 90%;
            border: 5px solid #FFB6C1;
        }
        h1 { color: #FF6347; margin-top: 0; }
        .input-group { margin-bottom: 15px; text-align: left; }
        label { color: #32CD32; font-weight: bold; display: block; margin-bottom: 5px;}
        input, select { 
            width: 100%; padding: 10px; border: 3px solid #FFD700; 
            border-radius: 10px; box-sizing: border-box; font-size: 1em;
        }
        .play-btn {
            background-color: #32CD32; color: white; border: none; padding: 15px;
            font-size: 1.5em; border-radius: 20px; cursor: pointer; width: 100%;
            font-weight: bold; box-shadow: 0 5px 0 #228B22; margin-top: 10px;
        }
        .play-btn:active { transform: translateY(5px); box-shadow: 0 0 0 #228B22; }
        #resultArea { display: none; margin-top: 20px; padding: 20px; background: #e0f7fa; border-radius: 15px; }
    </style>
</head>
<body>

    <div class="login-card">
        <h1>🎈 Kid World 🎈</h1>
        
        <form id="kidForm">
            <div class="input-group">
                <label>Kid's Name:</label>
                <input type="text" id="kidName" required>
            </div>
            <div class="input-group">
                <label>Age:</label>
                <select id="kidAge" required>
                    <option value="2">2</option>
                    <option value="3" selected>3</option>
                    <option value="4">4</option>
                    <option value="5">5+</option>
                </select>
            </div>
            <div class="input-group">
                <label>Mobile Number:</label>
                <input type="text" id="mobile" required>
            </div>
            <div class="input-group">
                <label>Town:</label>
                <input type="text" id="town" placeholder="e.g., Patna" required>
            </div>
            <div class="input-group">
                <label>State:</label>
                <select id="state" required>
                    <option value="Bihar">Bihar</option>
                    <option value="Maharashtra">Maharashtra</option>
                    <option value="Gujarat">Gujarat</option>
                    <option value="Delhi">Delhi</option>
                </select>
            </div>
            <button type="submit" class="play-btn">Let's Play! 🚀</button>
        </form>

        <div id="resultArea">
            <h2 id="welcomeMessage" style="color: #FF6347;"></h2>
            <p style="font-size: 1.2em;">Loading <strong id="languageText"></strong>...</p>
            <h3 id="moduleText" style="color: #32CD32;"></h3>
        </div>
    </div>

    <script>
        document.getElementById('kidForm').addEventListener('submit', async function(event) {
            event.preventDefault(); // Stops the page from refreshing

            // Gather the data from the form
            const profileData = {
                name: document.getElementById('kidName').value,
                age: parseInt(document.getElementById('kidAge').value),
                mobile: document.getElementById('mobile').value,
                town: document.getElementById('town').value,
                state: document.getElementById('state').value
            };

            try {
                // Send the data to your Go server
                const response = await fetch('http://localhost:8080/api/register', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(profileData)
                });

                const result = await response.json();

                // Hide the form and show the server's response!
                document.getElementById('kidForm').style.display = 'none';
                document.getElementById('resultArea').style.display = 'block';
                
                document.getElementById('welcomeMessage').innerText = result.message;
                document.getElementById('languageText').innerText = result.language;
                document.getElementById('moduleText').innerText = result.module;

            } catch (error) {
                alert("Could not connect to the server. Is your Go backend running?");
            }
        });
    </script>
</body>
</html>