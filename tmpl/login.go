package tmpl

import (
	"html/template"
	"net/http"
)

func RenderLogin(w http.ResponseWriter, render interface{}) {
	const html = `
<html lang="cn">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="陶士涵">
    <title>GO-IMAP网页版邮箱imap工具登录页</title>
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/element-ui@2.13.1/lib/theme-chalk/index.css">
    <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/element-ui@2.13.1/lib/index.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/jquery/dist/jquery.min.js"></script>
	<!-- Bootstrap core CSS -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
    <style>
        @media (min-width: 768px) {
            .bd-placeholder-img-lg {
                font-size: 3.5rem;
            }
        }
        html,
        body {
            height: 100%;
        }
        body {
            display: -ms-flexbox;
            display: flex;
            -ms-flex-align: center;
            align-items: center;
            padding-top: 40px;
            padding-bottom: 40px;
            background-color: #f5f5f5;
        }
        .form-signin {
            width: 100%;
            max-width: 400px;
            padding: 20px;
            margin: auto;
            background: #fff;
            -webkit-box-shadow: 0 1px 2px 0 rgba(101,129,156,.08);
            box-shadow: 0 1px 2px 0 rgba(101,129,156,.08);
        }
    </style>

</head>
<body class="text-center">
<div id="app" style="width:100%">
    <template>
        <el-form :model="ruleForm" :rules="rules" ref="ruleForm" class="form-signin">
          <h1 class="h3 mb-3 font-weight-normal">邮箱网页版IMAP工具</h1>
          <el-form-item  prop="server">
            <el-input v-model="ruleForm.server" placeholder="IMAP服务器如imap.sina.net:143"></el-input>
          </el-form-item>
          <el-form-item  prop="email">
            <el-input v-model="ruleForm.email" placeholder="邮箱地址"></el-input>
          </el-form-item>
          <el-form-item  prop="password">
            <el-input v-model="ruleForm.password" placeholder="密码"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button :loading="loading" type="primary" @click="submitForm('ruleForm')">立即登录</el-button>
            <el-button @click="resetForm('ruleForm')">重置</el-button>
          </el-form-item>
        </el-form>
        <p class="mt-5 mb-3 text-muted">&copy; 2020</p>
</template>           
</div>
</body>
<script>
	new Vue({
		el: '#app',
		data: {
			loading:false,
            ruleForm:{
                server:'',
                email:'',
                password:'',
            },
            rules: {
                server: [
                    { required: true, message: 'IMAP服务器如"imap.sina.net:143"包含端口号', trigger: 'blur' },
                ],
                email: [
                    { required: true, message: '邮箱地址', trigger: 'blur' },
                ],
                password: [
                    { required: true, message: '邮箱密码', trigger: 'blur' },
                ],
            },
		},
		methods: {
            //提交表单
			submitForm(formName){
                let _this=this;
                this.$refs[formName].validate((valid) => {
                  if (valid) {
                    var data={}
                    data.server=_this.ruleForm.server;
                    data.email=_this.ruleForm.email;
                    data.password=_this.ruleForm.password;
                    _this.loading=true;
                    $.post("/check",data,function(data){
                        if(data.code==200){
                            _this.$message({
                              message: data.msg,
                              type: 'success'
                            });
                            window.location.href="/";
                        }else{
                            _this.$message({
                              message: data.msg,
                              type: 'error'
                            });
                        }
                        _this.loading=false;
                    });
                  } else {
                    return false;
                  }
                });
			},
            //重置表单
            resetForm(formName) {
                this.loading=false;
                this.$refs[formName].resetFields();
            },
		}
	})

</script>
</html>
`
	t, _ := template.New("login").Parse(html)
	t.Execute(w, render)
}