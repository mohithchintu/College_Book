export function displaydate(date:Date|undefined)  {
    if (date === undefined) return "No date"
    const formattedDate = date.toLocaleString('en-US', { year: 'numeric', month: 'long', day: '2-digit' });
    return formattedDate;
}

