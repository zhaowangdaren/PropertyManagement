export default function (month) {
  if ((month % 3 ) > 0) {
    return parseInt(month / 3) + 1
  }
  return parseInt(month / 3)
}