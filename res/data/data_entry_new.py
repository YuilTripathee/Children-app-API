import json
import string

if __name__ == "__main__":
    with open("data/alphabets_raw.json", "r") as apl_fp:
        alp_raw_data = json.load(apl_fp)
        apl_fp.close()

    print("Data entry tool for alphabets:")
    print("==============================\n")

    # letters = ["a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"]
    letters = list(string.ascii_lowercase)
    alp_data = []
    id = 0
    for letter in letters:
        l_case = letter
        u_case = letter.upper()
        print("\n" + u_case)
        if letter is "a" or letter is "e" or letter is "i" or letter is "u":
            vowel = True
        else:
            vowel = False
        
        Wimg_ref = "http://127.0.0.1:8000/img/write/" + l_case
        img_ref = "http://127.0.0.1:8000/img/for/" + l_case
        soundN_ref = "http://127.0.0.1:8000/aud/norm/" + l_case
        soundF_ref = "http://127.0.0.1:8000/aud/for/" + l_case

        # adding data into an array
        alp_data_entity = { 
            "id" : id,
            "l_case" : l_case,
            "u_case" : u_case,
            "vowel_conf" : vowel,
            "write_img_ref" : Wimg_ref,
            "img_ref" : img_ref,
            "audio_for_ref" : soundF_ref,
            "audio_norm_ref" : soundN_ref
        }
        alp_data.append(alp_data_entity)
        id += 1

    with open("data/alphabets_local.json", "w", encoding="utf-8") as aplw_fp:
        json.dump(alp_data, aplw_fp, indent=4)
        print("\nData written into the json file")
        aplw_fp.close()
    