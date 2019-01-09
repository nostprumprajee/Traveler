package com.example.cosmeticapi;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class CosmeticController {
    List<Cosmetic> cosmetics = new ArrayList<>(Arrays.asList(
        new Cosmetic("LIP001", "XOXO Lipstick", 1),
        new Cosmetic("BRO001", "NYX Brush On Palette", 1)
    ));

@RequestMapping("/cosmetics")
    public List<Cosmetic> getAllCosmetics() {
        return cosmetics;
    }
}