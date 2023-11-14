resource "aws_route53_zone" "zone" {
  name = "miraikoji.dev"
}

resource "aws_route53_record" "www" {
  zone_id = aws_route53_zone.zone.zone_id
  name    = "www.miraikoji.dev"
  type    = "A"
  ttl     = "300"
  records = [aws_eip.eip.public_ip]
}
