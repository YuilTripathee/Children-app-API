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

        pre_img_ref = alp_raw_data[id]["img_ref"]
        # pre_img_ref = ''
        print("Existing image reference: " + pre_img_ref)
        img_ref = input("Enter the new image reference if necessary: ")
        if img_ref is '':
            img_ref = pre_img_ref
        
        pre_Wimg_ref = alp_raw_data[id]["write_img_ref"]
        # pre_Wimg_ref = ''
        print("Existing writing image reference: " + pre_Wimg_ref)
        Wimg_ref = input("Enter the new writing image reference if necessary: ")
        if Wimg_ref is '':
            Wimg_ref = pre_Wimg_ref

        # pre_aud_ref = alp_raw_data[id]["audio_ref"]
        # # pre_aud_ref = ''
        # print("Existing audio reference: " + pre_aud_ref)
        # sound_ref = input("Enter the new audio reference if necessary: ")
        # if sound_ref is '':
        #     sound_ref = pre_aud_ref
        
        # adding data into an array
        alp_data_entity = { 
            "id" : id,
            "l_case" : l_case,
            "u_case" : u_case,
            "vowel_conf" : vowel,
            "write_img_ref" : Wimg_ref,
            "img_ref" : img_ref
            # "audio_ref" : sound_ref
        }
        alp_data.append(alp_data_entity)
        id += 1

    with open("data/alphabets_raw.json", "w", encoding="utf-8") as aplw_fp:
        json.dump(alp_data, aplw_fp, indent=4)
        print("\nData written into the json file")
        aplw_fp.close()
    