# ssh-keygen -t rsa -f rickenbacker -N ''
# ssh -i ./rickenbacker ec2-user@18.183.232.33
resource "aws_key_pair" "key_pair" {
  key_name   = "rickenbacker"
  public_key = file("./rickenbacker.pub")
}

# Elastic IPを InternetGatewayに紐付ける
resource "aws_eip" "eip" {
  vpc        = true
  depends_on = [aws_internet_gateway.igw]
}
