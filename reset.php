<?php

    $token =  $_GET['token'];


?>
<!DOCTYPE html>
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- Add icon library-->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <style>
        body {
            font-family: Arial, Helvetica, sans-serif;
        }
        * {
            box-sizing: border-box;
        }
        .input-container {
            display: -ms-flexbox;
            /* IE10 */
            display: flex;
            width: 100%;
            margin-bottom: 15px;
        }
        .icon {
            padding: 10px;
            background: green;
            color: white;
            min-width: 50px;
            text-align: center;
        }
        .input-field {
            width: 100%;
            padding: 10px;
            outline: none;
        }
        .input-field:focus {
            border: 2px solid dodgerblue;
        }
        /* Set a style for the submit button */
        .btn {
            /* background-color: dodgerblue; */
            background-color: grey;
            color: white;
            padding: 15px 20px;
            border: none;
            cursor: pointer;
            width: 100%;
            opacity: 0.9;
        }
        .btn:hover {
            opacity: 1;
        }
        .fa-passwd-reset>.fa-lock {
            font-size: 0.85rem;
        }
    </style>
    <script>
        let check = function() {
            if (document.getElementById('password-1').value == document.getElementById('password-2').value) {
                document.getElementById("formSubmit").disabled = false;
                document.getElementById("formSubmit").style.background = 'blue';
                document.getElementById('message').style.color = 'green';
                document.getElementById('message').innerHTML = 'Password Matched';
            } else {
                document.getElementById("formSubmit").disabled = true;
                document.getElementById("formSubmit").style.background = 'grey';
                document.getElementById('message').style.color = 'red';
                document.getElementById('message').innerHTML = 'Password not matching';
            }
        }
        let validate = function() {
            console.log(document.getElementById('password-1').value)
            console.log(document.getElementById('password-2').value)
            if (document.getElementById('password-1').value.length < 5) {
                document.getElementById('pwd-length-1').innerHTML = 'Minimum 6 characters';
            } else {
                document.getElementById('pwd-length-1').innerHTML = '';
                check();
            }
            if (document.getElementById('password-2').value.length < 5) {
                document.getElementById('pwd-length-2').innerHTML = 'Minimum 6 characters';
            } else {
                document.getElementById('pwd-length-2').innerHTML = '';
                check();
            }
        }
    </script>
</head>
<body>
    <form action="/users/updatePassword" method="PUT" style="max-width:500px;margin:auto">
      <center>
      <img src="logo/payton.png">
      </center>
        <!-- Title  -->
        <center>
            <h2><span class="fa-passwd-reset fa-stack"><i class="fa fa-undo fa-stack-2x"></i><i class="fa fa-lock fa-stack-1x"></i></span>Reset your Password<span class="fa-passwd-reset fa-stack"><i class="fa fa-undo fa-stack-2x"></i><i class="fa fa-lock fa-stack-1x"></i></span></h2>
        </center>
        <!-- First Input Text  -->
        <div class="input-container"><i class="fa fa-key icon"></i>
            <input class="input-field" id="password-1" type="password" placeholder="Type your new password" name="password" oninput='validate();'>
        </div>
        <!-- Length  -->
        <span id="pwd-length-1"></span>
        <!-- Second Input Text  -->
        <div class="input-container"><i class="fa fa-key icon"></i>
            <input class="input-field" id="password-2" type="password" placeholder="Re-type your new password" name="confirmPassword" oninput='validate();'>
        </div>

        <div class="input-container" style="display: none;">
            <input class="input-field" id="token" type="text" name="token" value="<?php echo $token; ?>">
        </div>
        <!-- Length  -->
        <span id="pwd-length-2"></span>
        <!-- Validation Message - 1  -->
        <span id="message"></span>
        <button class="btn" id="formSubmit" type="submit" disabled>Register</button>
    </form>
</body>
</html>