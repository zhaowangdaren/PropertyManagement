export default function (pmkpi) {
  return (100 - pmkpi.YWNo * 0.5 - pmkpi.RWNo * 1 - pmkpi.IWNo * 3 - pmkpi.Other)
}