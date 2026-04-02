import csv 


'''
Just practicing opening and reading files as this simple syntax is usually 
done with copilot now and I might need it for interviews lol.
'''
def readCSV(filepath, output_file="suspcious_logins.txt"):
    with open(file=filepath, newline="") as f:
        reader = csv.DictReader(f)

        with open(output_file, "w", encoding="utf-8") as out:
            out.write("===Suspicous Login Report===\n")
            out.write("=" * 60 + "\n\n")

            count = 0

            for row in reader:
                suspicous = False
                reasoning = []

                if row["status"] == "FAILED":
                    suspicous = True
                    reasoning.append("user login failed")
                if row["home_country"] != "US":
                    suspicous = True
                    reasoning.append("user logged in outside home country")

                if suspicous:
                    line = f"Date: {row['timestamp']} | User: {row['user']} | Reason: {'. '.join(reasoning)}\n"
                    out.write(line)
                    count+=1
            out.write("\n" + "=" * 60 + "\n")
            out.write(f"Total suspicous entries: {count}")
                
def readwriteCSV(filepath, output_file="suspicous_logs.csv"):
    with open(filepath, newline="") as f:
        reader = csv.DictReader(f)

        with open(output_file, "w", encoding="utf-8") as out:
            fieldnames = ["timestamp", "user", "reason"]
            writer = csv.DictWriter(out, fieldnames=fieldnames)
            writer.writeheader()
            count = 0

            for row in reader:
                suspicous = False
                reasoning = []
                if row["status"] == "FAILED":
                    suspicous = True
                    reasoning.append("user failed login")
                if row["home_country"] != "US":
                    suspicous = True
                    reasoning.append("user login not from home country")
                
                if suspicous:
                    rowtowrite = {
                        "timestamp": row['timestamp'],
                        "user": row["user"],
                        "reason": ", ".join(reasoning)
                    }
                    writer.writerow(rowtowrite)
                    count+=1
            print(f"Suspicous logins written to txt and csv \nTotal suspicious entries found {count} ")

readCSV("Files/detectiondataset.csv")
readwriteCSV("Files/detectiondataset.csv")


def readCSV2txt(filepath, output_file="text.txt"):
    with open(filepath, newline="") as f:
        reader = csv.DictReader(f)

        with open(output_file, "w", encoding="utf-8") as out:
            out.write("===Suspicious Activity Report===\n")
            out.write("=" * 60 + "\n\n")


            count = 0
            for row in reader:

                suspicious = False
                reason = []
                if row['status'] == "FAILED":
                    suspicious = True
                    reason.append("user login failed")
                if row["home_country"] != "US":
                    suspicious = True
                    reason.append("log in not from home country")
                if suspicious:
                    write = f"Date: {row['timestamp']} | User: {row['user']} | Reason: {', '.join(reason)}\n"
                    out.write(write)
                    count+=1
            out.write("=" * 60 + "\n\n")
            out.write(f"Suspicious logins total: {count}")

def readcsvtocsv(filepath, output_file="text.csv"):
    with open(filepath, newline="") as f:
        reader = csv.DictReader(f)

        count = 0

        with open(output_file, "w", encoding="utf-8") as r:
            heading = ["timestamp", "user", "reason"]
            writer = csv.DictWriter(r, heading)

            for row in reader:
                
                suspicious = False
                reason = []
                if row['status'] == "FAILED":
                    suspicious = True
                    reason.append("user login failed")
                if row["home_country"] != "US":
                    suspicious = True
                    reason.append("log in not from home country")
                if suspicious:
                    write = {
                        "timestamp": row['timestamp'],
                        "user": row['user'],
                        "reason": ", ".join(reason)
                    }
                    writer.writerow(write)
                    count+=1
            print(f"Completed writing suspicious logins\nTotal: {count}")

readCSV2txt("Files/detectiondataset.csv")
readcsvtocsv("Files/detectiondataset.csv")