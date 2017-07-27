export default function (level) {
  switch (level) {
    case 1:
      return 'Urgent'//特急
    case 2:
      return 'Expedited'//加急
    case 3:
      return 'Worry'//急
  }
}