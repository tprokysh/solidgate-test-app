# solidgate-test-app
Hello! Thank you for this interesting test, and thank you for you'r attention! Let's start!
## Step-by-step set-up
1. First of all, of cource, git clone!
<pre>
cd ~/go/src
git clone https://github.com/tprokysh/solidgate-test-app.git
cd solidgate-test-app
</pre>
2.Copy config files, change data, and create database!
<pre>
cp etc/config.json.sample etc/config.json
cp etc/db.json.sapmle etc/db.json
</pre>
3.For callback you need install ngrok! <br>
https://dashboard.ngrok.com/get-started <br>
4.Time to run our server! <br>
For migration you need install gorm-goose:
<pre>
go get github.com/Altoros/gorm-goose/cmd/gorm-goose
</pre>
After that you need to put your database data into src/db/dbconf.yml <br>
So, after that you need run our migrations:
<pre>
cd src
~/go/bin/gorm-goose up
</pre>
I hope, migrate successfully done... <br>
So we can do that:
<pre>
cd ..
go run main.go
</pre>
5. So, server is up. For test our api, repository provides postman collection, import it into you'r postman, and start with create customer. <br>
<pre>
localhost:8080/customer
</pre>
6. Create some order for this customer <br>
<pre>
localhost:8080/order
</pre>
7. So try charge this order <br>
DONT FORGET PUT YOUR HGROK URL IN CALLBACK FOR SUCCESS CALLBACK!!!
<pre>
localhost:8080/customer/operation/charge
</pre>
If orderId still processing, change orderId in database with random number and try again <br>
8. After that we can try refund
<pre>
localhost:8080/customer/operation/refund
</pre>
If orderId still processing, change orderId in database with random number and try again <br>
9. Finally, try the recurring
<pre>
localhost:8080/customer/operation/recurring
</pre>
10. Taaa-daaaa, I hope all good, thank you again!
