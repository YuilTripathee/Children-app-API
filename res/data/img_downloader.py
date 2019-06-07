# the script downloads the images as in reference in JSON file and stores it in the project directories
import os
import json

if __name__ == "__main__":
    print("Image downloader software: ")

    # loading up image reference stored in JSON
    with open("data/alphabets_raw.json", "r") as alp_img_ref_fp:
        alp_img_ref_list = json.load(alp_img_ref_fp)
        alp_img_ref_fp.close()

    # img downloader
    for alphabet in alp_img_ref_list:
        alp_img_ref = alphabet["img_ref"]
        alp_img_w_ref = alphabet["write_img_ref"]
        img_cmd = "wget " + alp_img_ref + " -O " + "assets/img/for_img/" + alphabet["l_case"] + ".png"
        img_w_cmd = "wget " + alp_img_w_ref + " -O " + "assets/img/writer_img/" + alphabet["u_case"] + ".gif"
        try:
            os.system(img_cmd)
            os.system(img_w_cmd)
        except:
            next

    print("All the files from references are downloaded.") 