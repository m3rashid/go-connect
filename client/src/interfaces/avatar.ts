export interface IAvatar {
  sex: "man" | "woman";
  earSize: "small" | "big";
  hairStyle: "normal" | "thick" | "mohawk" | "womanLong" | "womanShort";
  hatStyle: "none" | "beanie" | "turban";
  glassesStyle: "none" | "round" | "square";
  noseStyle: "short" | "long" | "round";
  mouthStyle: "laugh" | "smile" | "peace";
  shirtStyle: "hoody" | "short" | "polo";
  faceColor: string;
  hairColor: string;
  hatColor: string;
  shirtColor: string;
  bgColor: string;
  avatarID: string;
  isGradient?: boolean;
}
