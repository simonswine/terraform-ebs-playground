variable "instances" {
  default = 1
}

variable "zone" {
  default = 1
}

resource "aws_ebs_volume" "volume" {
  count             = "${var.instances}"
  availability_zone = "${var.zone}"
  size              = 40

  encrypted = "false"

  tags {
    Name = "HelloWorld2"
  }
}
